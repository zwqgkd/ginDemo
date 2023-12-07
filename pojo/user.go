package pojo

import ("golangAPI/database"
"log"

)

type User struct { //
	Id       int    `json:"UserId" binding:"required"`
	Name     string `json:"UserName" binding:"required"`
	Password string `json:"UserPassword" binding:"required,min=4,max=20,userpasd"`
	Email    string `json:"UserEmail" binding:"email"`
}

type Users struct{
	UserList []User `json:"UserList" binding:"required,gt=0,lt=3"`
	UserListSize int `json:"UserListSize"`
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


//CheckUserPassword
func CheckUserPassword(name string,password string) User{
	user:=User{}
	database.DBConnect.Where("name=? and password=?",name,password).First(&user)
	return user
}