package validate

import (
	"fmt"
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

//validate1 instance create func 创建验证实体函数
func (s *StructData) Create(err ...error) *Validation {
	return nil
}
func (s *StructData) Validation(err ...error) *Validation {
	v := NewValidation(s)
	if len(err) > 0 && err[0] != nil {
		return v.WithErr(err[0])
	}

	// collect field filter/validate1 rules from struct tags
	s.parseRulesFromTag(v)

	return v
}
func (s *StructData) parseRulesFromTag(v *Validation) {
	if s.ValidateTag == "" {
		s.ValidateTag = globalOpt.ValidateTag
	}
	if s.FilterTag == "" {
		s.FilterTag = globalOpt.FilterTag
	}
	vt := s.valueTpy
	for i := 0; i < vt.NumField(); i++ {
		name := vt.Field(i).Name
		if name[0] >= 'a' && name[0] <= 'z' {
			continue
		}
		s.fieldNames[name] = 1
		vRule := vt.Field(i).Tag.Get(s.ValidateTag)
		if vRule != "" {
			fmt.Println(vRule)
		}
		fRule := vt.Field(i).Tag.Get(s.FilterTag)
		if fRule != ""{
			fmt.Println(fRule)
		}
	}
}
