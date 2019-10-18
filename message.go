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

/*************************************************************
 * validators messages
 *************************************************************/

// default internal error message for some rules.
var defMessages = map[string]string{
	"_":       "{field} did not pass validate", // default message
	"_filter": "{field} data is invalid",       // data filter error
	// int value
	"min": "{field} min value is %d",
	"max": "{field} max value is %d",
	// type check: int
	"isInt":  "{field} value must be an integer",
	"isInt1": "{field} value must be an integer and mix value is %d",      // has min check
	"isInt2": "{field} value must be an integer and in the range %d - %d", // has min, max check
	"isInts": "{field} value must be an int slice",
	"isUint": "{field} value must be an unsigned integer(>= 0)",
	// type check: string
	"isString":  "{field} value must be an string",
	"isString1": "{field} value must be an string and min length is %d", // has min len check
	// length
	"minLength": "{field} min length is %d",
	"maxLength": "{field} max length is %d",
	// string length. calc rune
	"stringLength":  "{field} length must be in the range %d - %d",
	"stringLength1": "{field} min length is %d",
	"stringLength2": "{field} length must be in the range %d - %d",

	"isURL":     "{field} must be an valid URL address",
	"isFullURL": "{field} must be an valid full URL address",

	"isFile":  "{field} must be an uploaded file",
	"isImage": "{field} must be an uploaded image file",

	"enum":  "{field} value must be in the enum %v",
	"range": "{field} value must be in the range %d - %d",
	// required
	"required": "{field} is required",
	// field compare
	"eqField":  "{field} value must be equal the field %s",
	"neField":  "{field} value cannot be equal the field %s",
	"ltField":  "{field} value should be less than the field %s",
	"lteField": "{field} value should be less than or equal to field %s",
	"gtField":  "{field} value must be greater the field %s",
	"gteField": "{field} value should be greater or equal to field %s",
}

// Translator definition
type Translator struct {
	// field map {"field name": "display name"}
	fieldMap map[string]string
	// message data map
	messages map[string]string
}

func NewTranslator() *Translator {
	defMap := make(map[string]string)
	for k, v := range defMessages {
		defMap[k] = v
	}
	fidldMap := make(map[string]string)
	return &Translator{
		fieldMap: fidldMap,
		messages: defMap,
	}

}
