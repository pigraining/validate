package validate

import "reflect"

var globalOpt = &GlobalOption{
	StopOnError: true,
	SkipOnEmpty: true,
	// tag name in struct tags
	FilterTag: filterTag,
	// tag name in struct tags
	ValidateTag: validateTag,
}

/*
******************************************
		CreateStruct
******************************************
*/

//Create Struct instance validate 创建结构体的validate实体
func CreateStruct(s interface{},sence ...string) *Validation{
	return NewWithError(FromStruct(s)).SetCence()
}

// FromStruct create a Data from struct
func FromStruct(s interface{}) (*StructData, error) {
	data := &StructData{
		ValidateTag: globalOpt.ValidateTag,
		// init map
		fieldNames:  make(map[string]int),
		fieldValues: make(map[string]reflect.Value),
	}

	if s == nil {
		return data, ErrInvalidData
	}

	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	typ := val.Type()
	if val.Kind() != reflect.Struct || typ == timeType {
		return data, ErrInvalidData
	}

	data.src = s
	data.value = val
	data.valueTpy = typ

	return data, nil
}