package rest

import (
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"net/http"
	"reflect"
)

type RequestOptions func(*Options)

type Options struct {
	Header map[string]string
}

var defaultOpts = Options{Header: map[string]string{
	httputil.HeaderContent: httputil.JsonHeaderContent,
}}

func Get(url string, dest interface{}, opts ...RequestOptions) aerror.Error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return aerror.NewErrorf(nil, aerror.Code.CParamsError, "type of 'dest' is must a ptr")
	}

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.OtherNetworkError, "new request failed")
	}

	reqOpts := defaultOpts
	for _, o := range opts {
		o(&reqOpts)
	}

	// set request header
	setHeader(req, reqOpts)

	resp, err := client.Do(req)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.OtherNetworkError, "request failed")
	}

	err = httputil.DecodeResponseBody(resp, &dest)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SRequestBodyDecodeFailed, "decode response body failed")
	}

	return nil
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func setHeader(req *http.Request, opts Options) {
	for k, v := range opts.Header {
		req.Header.Add(k, v)
	}
}
