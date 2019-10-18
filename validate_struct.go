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




