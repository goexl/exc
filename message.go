package exc

// MessageException 带消息的异常
type MessageException interface {
	error

	// Message 返回错误消息
	Message() string
}
