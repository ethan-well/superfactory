package rest

import (
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
	"reflect"
)

func Get(url string, dest interface{}) *xerror.Error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "type of 'dest' is must a ptr")
	}
	
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.OtherNetworkError, "new request failed")
	}
	req.Header.Add(httputil.HeaderContent, httputil.JsonHeaderContent)
	
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

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
