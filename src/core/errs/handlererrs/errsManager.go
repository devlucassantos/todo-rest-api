package handlererrs

import (
	"strings"
	"todo/src/core/errs/handlererrs/msgs"
)

func GetAuthBindError(err error) (string, InvalidFields) {
	var errMessage string
	invalidFields := InvalidFields{}
	if strings.Contains(err.Error(), "name") {
		errMessage = "Invalid Name"
		invalidFields.AppendField("name", msgs.CheckDataType)
	} else if strings.Contains(err.Error(), "email") {
		errMessage = "Invalid Email"
		invalidFields.AppendField("email", msgs.CheckDataType)
	} else if strings.Contains(err.Error(), "password") {
		errMessage = "Invalid Password"
		invalidFields.AppendField("password", msgs.CheckDataType)
	}

	return errMessage, invalidFields
}
