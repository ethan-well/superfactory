package rest

import (
	"bytes"
	"encoding/json"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
	"reflect"
)

func Post(url string, body interface{}, dest interface{}) *xerror.Error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "type of 'dest' is must a ptr")
	}
	
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	
	date, err := json.Marshal(body)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.BUnexpectedData, "json marshal failed")
	}
	
	b := bytes.NewBuffer(date)
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.OtherNetworkError, "new request failed")
	}
	req.Header.Add(httputil.HeaderContent, httputil.JsonHeaderContent)
	req.Header.Add(httputil.HeaderAccept, httputil.JsonHeaderAccept)
	
	resp, err := client.Do(req)
	if err != nil {
		return xerror.NewErrorf(nil, xerror.Code.OtherNetworkError, "request failed")
	}
	
	err = httputil.DecodeResponseBody(resp, &dest)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SRequestBodyDecodeFailed, "decode response body failed")
	}
	
	return nil
}