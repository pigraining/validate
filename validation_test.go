package validate

import (
	"fmt"
	"testing"
	"time"
)

// UserForm struct
type UserForm struct {
	Name     string    `validate:"required|minLen:7"`
	Email    string    `validate:"email"`
	Age      int       `validate:"required|int|min:1|max:99"`
	CreateAt int       `validate:"min:1"`
	Safe     int       `validate:"-"`
	UpdateAt time.Time `validate:"required"`
	Code     string    `validate:"customValidator"`
}

// CustomValidator custom validator in the source struct.
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

// Messages you can custom validator error messages.
func (f UserForm) Messages() map[string]string {
	return MS{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}

// Translates you can custom field translates.
func (f UserForm) Translates() map[string]string {
	return MS{
		"Name":  "User Name",
		"Email": "User Email",
	}
}

func TestCreateStruct(t *testing.T) {
	u := &UserForm{
		Name: "inhere",
	}

	v:= CreateStruct(u)
	// v := validate1.New(u)
	fmt.Println(v)
	/*if v.Validate() { // validate1 ok
		// do something ...
	} else {
		fmt.Println(v.Errors) // all error messages
		fmt.Println(v.Errors.One()) // returns a random error message text
		fmt.Println(v.Errors.Field("Name")) // returns error messages of the field
	}*/
}
