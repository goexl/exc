package exc

import (
	`encoding/json`
	`fmt`
	`strings`

	`github.com/goexl/gox`
)

var (
	_           = NewException
	_ Exception = (*exceptionDefault)(nil)
)

type exceptionDefault struct {
	code    int
	message string
	fields  gox.Fields
}

// NewException 创建异常
func NewException(code int, message string, fields ...gox.Field) *exceptionDefault {
	return &exceptionDefault{
		code:    code,
		message: message,
		fields:  fields,
	}
}

func (e *exceptionDefault) Code() int {
	return e.code
}

func (e *exceptionDefault) Message() string {
	return e.message
}

func (e *exceptionDefault) Fields() gox.Fields {
	return e.fields
}

func (e *exceptionDefault) MarshalJSON() (bytes []byte, err error) {
	output := make(map[string]interface{})
	output[`code`] = e.code
	output[`message`] = e.message

	if 0 < len(e.fields) {
		data := make(map[string]interface{})
		for _, field := range e.fields {
			data[field.Key()] = field.Value()
		}
		output[`data`] = data
	}
	bytes, err = json.Marshal(output)

	return
}

func (e *exceptionDefault) Error() (str string) {
	if bytes, err := e.MarshalJSON(); nil != err {
		str = e.error()
	} else {
		str = string(bytes)
	}

	return
}

func (e *exceptionDefault) error() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf(`code = %d, `, e.code))
	sb.WriteString(fmt.Sprintf(`message = %s, `, e.message))
	sb.WriteString(fmt.Sprintf(`data = %s`, e.fields.String()))
	sb.WriteRune('}')

	return sb.String()
}
