package xerror

type ErrorCode struct {
	CParamsError    string
	CreateUserError string
	// 客户端传递的非预期参数
	CUnexpectRequestDate string

	SGetLockeErr string

	SSqlExecuteErr   string
	SUnexpectedErr   string
	SRedisExecuteErr string
	SFileExisted     string

	SUnimplementedInterface string

	BDataIsNotAllow          string
	BUnexpectedBlankVariable string
	BUnexpectedData          string
	BCreateFileFailed        string

	OtherSystemError  string
	OtherNetworkError string
}

var Code = ErrorCode{
	CParamsError:         "C0001",
	CreateUserError:      "C0002",
	CUnexpectRequestDate: "C0003",

	SGetLockeErr:            "S0001",
	SSqlExecuteErr:          "S0002",
	SUnexpectedErr:          "S0003",
	SRedisExecuteErr:        "S0004",
	SFileExisted:            "S0005",
	SUnimplementedInterface: "S0006",

	BDataIsNotAllow: "B0001",

	BUnexpectedBlankVariable: "B0002",

	BUnexpectedData: "B0003",

	BCreateFileFailed: "B0004",

	// 其他服务错误
	OtherSystemError:  "O0001",
	OtherNetworkError: "O0002",
}
