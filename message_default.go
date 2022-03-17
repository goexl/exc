package exc

import (
	`fmt`
)

var (
	_                  = NewMessage
	_ MessageException = (*messageDefault)(nil)
)

type messageDefault struct {
	message string
}

// NewMessage 创建带消息的异常
func NewMessage(message string) *messageDefault {
	return &messageDefault{
		message: message,
	}
}

func (m *messageDefault) Message() string {
	return m.message
}

func (m messageDefault) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{message: %s}`, m.message)), nil
}

func (m *messageDefault) Error() string {
	return fmt.Sprintf(`message = %s, `, m.message)
}
