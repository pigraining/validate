package validate

import "reflect"

// some default value settings.
const (
	filterTag     = "filter"
	filterError   = "_filter"
	validateTag   = "validate"
	validateError = "_validate"
	// sniff Length, use for detect file mime type
	sniffLen = 512
	// 32 MB
	defaultMaxMemory int64 = 32 << 20
)

// M is short name for map[string]interface{}
type M map[string]interface{}

// MS is short name for map[string]string
type MS map[string]string

// SValues simple values
type SValues map[string][]string

type Validation struct {
	// source input data
	data DataFace
	// all validated fields list
	// fields []string
	// filtered/validated safe data
	safeData M
	// filtered clean data
	filteredData M
	// Errors for the validate
	Errors Errors
	// CacheKey for cache rules
	// CacheKey string
	// StopOnError If true: An error occurs, it will cease to continue to verify
	StopOnError bool
	// SkipOnEmpty Skip check on field not exist or value is empty
	SkipOnEmpty bool
	// UpdateSource Whether to update source field value, useful for type validate
	UpdateSource bool
	// CheckDefault Whether to validate the default value set by the user
	CheckDefault bool
	// CachingRules switch. default is False
	// CachingRules bool
	// save user set default values
	defValues map[string]interface{}
	// mark has error occurs
	hasError bool
	// mark is filtered
	hasFiltered bool
	// mark is validated
	hasValidated bool
	// validate rules for the validation
	rules []*Rule
	// validators for the validation
	validators map[string]int
	// validator func meta info
	validatorMetas map[string]*funcMeta
	// validator func reflect.Value map
	validatorValues map[string]reflect.Value
	// translator instance
	trans *Translator
	// current scene name
	scene string
	// scenes config.
	// {
	// 	"create": {"field0", "field1"}
	// 	"update": {"field0", "field2"}
	// }
	scenes SValues
	// should checked fields in current scene.
	sceneFields map[string]uint8
	// filtering rules for the validation
	filterRules []*FilterRule
	// filter func reflect.Value map
	filterValues map[string]reflect.Value
}

func NewEmptyValidate(scean ...string)*Validation{
	return NewValidation(nil,scean...)
}

func NewValidation(data DataFace,scean ...string)*Validation{
	v := &Validation{
		data:data,
		Errors:make(Errors),
		trans:NewTranslator(),
	}
	v.validatorValues = map[string]reflect.Value{
		//"required":reflect.ValueOf()
	}
	v.SetCence(scean...)
	return v
}

//使用这个方法是为了接收并处理err
func NewWithError(data DataFace,err error)*Validation{
	if data == nil{
		if err != nil{
			return NewValidation(data).WithErr(err)
		}
		return NewValidation(data)
	}
	return data.Validation()
}

func(v *Validation)SetCence(scene...string)*Validation{
	if len(scene)>0{
		v.scene = scene[0]
	}
	return v
}

func(v *Validation)WithErr(err error)*Validation{
	if err != nil{
		v.AddErr(validateError,err.Error())
	}
	return v
}

func(v *Validation)AddErr(filed,msg string){
	if !v.hasError{
		v.hasError = true
	}
	v.Errors.Add(filed,msg)
}