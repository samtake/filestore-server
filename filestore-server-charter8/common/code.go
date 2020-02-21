package common

// ErrorCode : 错误码
type ErrorCode int32

const (
	_ int32 = iota + 9999
	// StatusOK : 10000 正常
	StatusOK
	// StatusParamInvalid :  10001 请求参数无效
	StatusParamInvalid
	// StatusServerError : 10002 服务出错
	StatusServerError
	// StatusRegisterFailed : 10003 注册失败
	StatusRegisterFailed
	// StatusLoginFailed : 10004 登录失败
	StatusLoginFailed
	// StatusInvalidToken : 10005 token无效
	StatusInvalidToken
)
