package validate

import "reflect"

// FilterRule definition
type FilterRule struct {
	// fields to filter
	fields []string
	// filter name list
	filters []string
	// filter args. { index: "args" }
	filterArgs map[int]string
}

// Rule definition
type Rule struct {
	// eg "create" "update"
	scene string
	// need validate fields. allow multi.
	fields []string
	// is optional, only validate on value is not empty. sometimes
	optional bool
	// skip validate not exist field/empty value
	skipEmpty bool
	// default value setting
	defValue interface{}
	// error message
	message string
	// error messages, if fields contains multi field.
	// eg {
	// 	"field": "error message",
	// 	"field.validator": "error message",
	// }
	messages map[string]string
	// validator name, allow multi validators. eg "min", "range", "required"
	validator string
	// arguments for the validator
	arguments []interface{}
	// --- some hooks function
	// has beforeFunc. if return false, skip validate current rule
	beforeFunc func(field string, v *Validation) bool // func (val interface{}) bool
	// you can custom filter func
	filterFunc func(val interface{}) (interface{}, error)
	// custom check func's mate info
	checkFuncMeta *funcMeta
	// custom check is empty.
	emptyChecker func(val interface{}) bool
}

type funcMeta struct {
	fv   reflect.Value
	name string
	// readonly cache
	numIn  int
	numOut int
	// is internal built in validator
	isInternal bool
	// last arg is like "... interface{}"
	isVariadic bool
}




// Translator definition
type Translator struct {
	// field map {"field name": "display name"}
	fieldMap map[string]string
	// message data map
	messages map[string]string
}




