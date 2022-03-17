package exc

import (
	`fmt`
)

var (
	_      = NewCode
	_ Code = (*codeDefault)(nil)
)

type codeDefault struct {
	code int
}

// NewCode 创建带状态码的异常
func NewCode(code int) *codeDefault {
	return &codeDefault{
		code: code,
	}
}

func (c *codeDefault) Code() int {
	return c.code
}

func (c codeDefault) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{code: %d}`, c.code)), nil
}

func (c *codeDefault) Error() string {
	return fmt.Sprintf(`code = %d, `, c.code)
}
