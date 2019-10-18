package validate

import "strings"

type Rules []*Rule

/*
************************************************
				Rule Struct
************************************************
*/

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

func NewRule(fields, validator string, args ...interface{}) *Rule {
	return &Rule{
		fields:    strings.Split(fields, ","),
		arguments: args,
		validator: validator,
	}
}
/*************************************************************
 * add validate rules
 *************************************************************/

// StringRule add field rules by string
// Usage:
// 	v.StringRule("name", "required|string|minLen:6")
// 	// will try convert to int before apply validate.
// 	v.StringRule("age", "required|int|min:12", "toInt")
func (v *Validation) StringRule(filed, rule string, filterRule ...string) *Validation {
	rules := strings.Split(rule, "|")
	for _,validator := range rules{
		if strings.Contains(validator,":"){

		}else {

		}
	}
	return nil
}


func(v *Validation)AddRule(fields, validator string, args ...interface{})*Rule{
	rule := NewRule(fields,validator,args...)
	v.rules = append(v.rules,rule)
	return rule
}


