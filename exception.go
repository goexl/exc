package exc

// Exception 异常，符合条件的错误统一处理
type Exception interface {
	Code
	MessageException
	Fields
}
