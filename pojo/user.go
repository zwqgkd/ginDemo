package pojo

import ("golangAPI/database"
"log"

)

type User struct { //
	Id       int    `json:"UserId" validate:"required"`
	Name     string `json:"UserName" validate:"required"`
	Password string `json:"UserPassword" validate:"required,min=4,max=20"`
	Email    string `json:"UserEmail" validate:"email"`
}

func FindAllUsers() []User {
	var users []User
	database.DBConnect.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	var user User
	//gorm会进行一种隐式的类型转换，尝试转为数据库里面的类型
	database.DBConnect.Where("id=?", userId).First(&user)
	return user
}

//post
func CreateUser(user User) User{
	database.DBConnect.Create(&user)
	return user
}

//delete
func DeleteUser(userId string) bool{
	user:=User{}
	tx:=database.DBConnect.Delete(user,userId)
	//tx:=database.DBConnect.Where("id=?",userId).Delete(&user)
	if tx.Error!=nil{
		log.Println("Error:",tx.Error.Error())
		return false
	}
	return true
}

//updateUser
func UpdateUser(userId string, user User) User{
	database.DBConnect.Where("id=?",userId).Updates(user)
	return user
}
