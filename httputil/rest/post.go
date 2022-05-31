package rest

import (
	"bytes"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"net/http"
	"reflect"
)

func Post(url string, body interface{}, dest interface{}, opts ...RequestOptions) (*http.Response, aerror.Error) {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "type of 'dest' is must a ptr")
	}

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	date, err := json.Marshal(body)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "json marshal failed")
	}

	b := bytes.NewBuffer(date)
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.OtherNetworkError, "new request failed")
	}

	reqOpts := defaultOpts
	for _, o := range opts {
		o(&reqOpts)
	}

	// set request header
	setHeader(req, reqOpts)

	resp, err := client.Do(req)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.OtherNetworkError, "request failed")
	}

	err = httputil.DecodeResponseBody(resp, &dest)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SRequestBodyDecodeFailed, "decode response body failed")
	}

	return resp, nil
}
