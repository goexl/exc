package exc

import (
	`encoding/json`
	`fmt`
	`strings`

	`github.com/goexl/gox`
)

var (
	_       = NewField
	_ Field = (*fieldDefault)(nil)
)

type fieldDefault struct {
	message string
	field   gox.Field
}

// NewField 创建带字段的异常
func NewField(message string, field gox.Field) *fieldDefault {
	return &fieldDefault{
		message: message,
		field:   field,
	}
}

func (f *fieldDefault) Message() string {
	return f.message
}

func (f *fieldDefault) Field() gox.Field {
	return f.field
}

func (f *fieldDefault) MarshalJSON() (bytes []byte, err error) {
	output := make(map[string]interface{})
	output[`message`] = f.message
	output[f.Field().Key()] = f.Field().Value()
	bytes, err = json.Marshal(output)

	return
}

func (f *fieldDefault) Error() (str string) {
	if bytes, err := f.MarshalJSON(); nil != err {
		str = f.error()
	} else {
		str = string(bytes)
	}

	return
}

func (f *fieldDefault) error() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf(`message = %s, `, f.message))
	sb.WriteString(fmt.Sprintf(`%s = %v`, f.field.Key(), f.field.Value()))
	sb.WriteRune('}')

	return sb.String()
}
