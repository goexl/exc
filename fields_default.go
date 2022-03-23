package exc

import (
	`encoding/json`
	`fmt`
	`strings`

	`github.com/goexl/gox`
)

var (
	_        = NewFields
	_ Fields = (*fieldsDefault)(nil)
)

type fieldsDefault struct {
	message string
	fields  gox.Fields
}

// NewFields 创建带多个字段列表的异常
func NewFields(message string, fields ...gox.Field) *fieldsDefault {
	return &fieldsDefault{
		message: message,
		fields:  fields,
	}
}

func (f *fieldsDefault) Message() string {
	return f.message
}

func (f *fieldsDefault) Fields() gox.Fields {
	return f.fields
}

func (f *fieldsDefault) MarshalJSON() (bytes []byte, err error) {
	output := make(map[string]interface{})
	output[`message`] = f.message

	if 0 < len(f.fields) {
		data := make(map[string]interface{})
		for _, field := range f.fields {
			data[field.Key()] = field.Value()
		}
		output[`data`] = data
	}
	bytes, err = json.Marshal(output)

	return
}

func (f *fieldsDefault) Error() (str string) {
	if bytes, err := f.MarshalJSON(); nil != err {
		str = f.error()
	} else {
		str = string(bytes)
	}

	return
}

func (f *fieldsDefault) error() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf(`message = %s, `, f.message))
	sb.WriteString(fmt.Sprintf(`data = %s`, f.fields.String()))
	sb.WriteRune('}')

	return sb.String()
}
