package exc

// Code 带状态码的异常
type Code interface {
	error

	// Code 返回错误码
	Code() int
}
