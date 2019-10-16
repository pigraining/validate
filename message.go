package validate

import "errors"

var (
	ErrSetValue = errors.New("set value failure")
	// ErrNoData = errors.New("validate: no any data can be collected")
	ErrNoField     = errors.New("field not exist in the source data")
	ErrEmptyData   = errors.New("please input data use for validate")
	ErrInvalidData = errors.New("invalid input data")
)
/*
***************************************************
**                Errors Struct                  **
***************************************************
*/

type Errors map[string][]string

// Add a error for the field 给这个类型添加error
func (es Errors) Add(field, msg string) {
	if ss, ok := es[field]; ok {
		es[field] = append(ss, msg)
	} else {
		es[field] = []string{msg}
	}
}
