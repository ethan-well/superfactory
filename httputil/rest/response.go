package rest

import (
	"encoding/json"
	"github.com/ItsWewin/superfactory/httputil"
	"net/http"
)

type Response struct {
	Result    interface{} `json:"result,omitempty"`
	Succeed   bool        `json:"succeed"`
	ErrorInfo *ErrorInfo  `json:"error_info,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func SucceedJsonResponse(w http.ResponseWriter, response interface{}) (int, error) {
	res := Response{
		Result:  response,
		Succeed: true,
	}

	bt, err := json.Marshal(res)
	if err != nil {
		return 0, err
	}

	// set response header
	// Content-Type => application/json
	w.Header().Set(httputil.HeaderContent, httputil.JsonHeaderContent)
	return w.Write(bt)
}

func FailJsonResponse(w http.ResponseWriter, code, message string) (int, error) {
	res := Response{
		Succeed: false,
		ErrorInfo: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}

	w.Header().Set(httputil.HeaderContent, httputil.JsonHeaderContent)
	bt, err := json.Marshal(res)
	if err != nil {
		return 0, err
	}

	// set response header
	// Content-Type => application/json
	w.Header().Set(httputil.HeaderContent, httputil.JsonHeaderContent)
	return w.Write(bt)
}
