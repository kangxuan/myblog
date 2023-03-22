package util

import (
	"github.com/astaxie/beego/validation"
)

func GetValidationMessage(valid validation.Validation) (hasErrors bool, message string) {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			message += err.Key + ":" + err.Message + ";"
		}

		return true, message
	} else {
		return false, ""
	}
}
