package xerror

type ErrorMessage struct {
	ParamsError string
	SqlExecErr  string
	CForbidden  string
}

var Message = ErrorMessage{
	ParamsError: "params is invalid",
	SqlExecErr:  "sql exec error",
	CForbidden:  "Forbidden",
}
