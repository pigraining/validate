package validate

import (
	"reflect"
	"time"
)

// DataFace interface definition
type DataFace interface {
	Type() uint8
	Get(key string) (interface{}, bool)
	Set(field string, val interface{}) error
	// validation instance create func
	Create(err ...error) *Validation
	Validation(err ...error) *Validation
}

var timeType = reflect.TypeOf(time.Time{})

/*************************************************************
 * Struct Data
 *************************************************************/

// StructData definition
type StructData struct {
	// source struct data, from user setting
	src interface{}
	// from reflect source Struct
	value reflect.Value
	// source struct reflect.Type
	valueTpy reflect.Type
	// field names in the src struct
	fieldNames map[string]int
	// field values cache
	fieldValues map[string]reflect.Value
	// FilterTag name in the struct tags.
	FilterTag string
	// ValidateTag name in the struct tags.
	ValidateTag string
}

func (s *StructData) Type() uint8 {
	return 0
}
func (s *StructData) Get(Key string) (interface{}, bool) {
	return nil, true
}
func (s *StructData) Set(filed string, val interface{}) error {
	return nil
}

//validate instance create func 创建验证实体函数
func (s *StructData) Create(err ...error) *Validation {
	return nil
}
func (s *StructData) Validation(err ...error) *Validation {
	return nil
}
