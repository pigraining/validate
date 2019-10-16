package validate

/*
***********************************************
			GlobalOption Struct
***********************************************
*/

// GlobalOption settings for validate
type GlobalOption struct {
	// FilterTag name in the type tags.
	FilterTag string
	// ValidateTag in the type tags.
	ValidateTag string
	// StopOnError If true: An error occurs, it will cease to continue to verify
	StopOnError bool
	// SkipOnEmpty Skip check on field not exist or value is empty
	SkipOnEmpty bool
	// CheckDefault whether validate the default value
	CheckDefault bool
}
