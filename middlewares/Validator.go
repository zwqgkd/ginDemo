package middlewares

import (
	"regexp"
	"golangAPI/pojo"
	"github.com/go-playground/validator/v10"

)

func UserPasd(field validator.FieldLevel) bool {
	if matched,_:=regexp.MatchString(`^[a-zA-Z0-9]{4,20}$`, field.Field().String());matched{
		return true
	}
	return false
}

func UserList(field validator.StructLevel){
	users:=field.Current().Interface().(pojo.Users)
	if users.UserListSize==len(users.UserList){

	}else{
		field.ReportError(users.UserListSize,"UserListSize","Users","UserListSizeMustEqualUserList","")
	}
}
