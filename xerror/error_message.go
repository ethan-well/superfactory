package xerror

type ErrorMessage struct {
	ParamsError string
	SqlExecErr  string
}

var Message = ErrorMessage{
	ParamsError: "params is invalid",
	SqlExecErr:  "sql exec error",
}
