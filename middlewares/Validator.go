package middlewares

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func UserPasd(field validator.FieldLevel) bool {
	
	if matched,_:=regexp.MatchString(`^[a-zA-Z0-9]{4,20}$`, field.Field().String());matched{
		return true
	}
	return false
}