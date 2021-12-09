package rest

import (
	"bytes"
	"encoding/json"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
	"reflect"
)

func Post(url string, body interface{}, dest interface{}, opts ...RequestOptions) (*http.Response, *xerror.Error) {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "type of 'dest' is must a ptr")
	}

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	date, err := json.Marshal(body)
	if err != nil {
		return nil, xerror.NewErrorf(err, xerror.Code.BUnexpectedData, "json marshal failed")
	}

	b := bytes.NewBuffer(date)
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, xerror.NewErrorf(err, xerror.Code.OtherNetworkError, "new request failed")
	}

	reqOpts := defaultOpts
	for _, o := range opts {
		o(&reqOpts)
	}

	// set request header
	setHeader(req, reqOpts)

	resp, err := client.Do(req)
	if err != nil {
		return nil, xerror.NewErrorf(nil, xerror.Code.OtherNetworkError, "request failed")
	}

	err = httputil.DecodeResponseBody(resp, &dest)
	if err != nil {
		return nil, xerror.NewErrorf(err, xerror.Code.SRequestBodyDecodeFailed, "decode response body failed")
	}

	return resp, nil
}
