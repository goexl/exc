package exc

import (
	`github.com/goexl/gox`
)

// Fields 带多个字段列表的异常
type Fields interface {
	error
	MessageException

	// Fields 返回错误实体
	// 在某些错误下，可能需要返回额外的信息给前端处理
	// 比如，认证错误，需要返回哪些字段有错误
	Fields() gox.Fields
}
