package pojo

import (
	"context"
	"golangAPI/database"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//MongoDB
func MgoCreateUser(user User)User{
	database.MgoCollection.InsertOne(context.TODO(),user)
	return user
}

func MgoFindAllUsers() []User{
	var users []User
	cur,err:=database.MgoCollection.Find(context.TODO(),bson.M{})
	if err!=nil{
		log.Println(err)
	}
	for cur.Next(context.TODO()){
		var user User	
		err:=cur.Decode(&user)
		if err!=nil{
			log.Println(err)
		}
		users=append(users,user)
	}
	return users
}

func MgoFindById(id string) User{
	userId,_:=strconv.Atoi(id)
	user := User{}
	database.MgoCollection.FindOne(context.TODO(),bson.D{{"id",userId}}).Decode(&user)
	return user
}

func MgoPutUser(id string, user User) User{
	userId,_:=strconv.Atoi(id)

	updateUserId:=bson.D{{"id",userId}}
	updateData:=bson.D{{"$set",user}}
	opt:=options.Update().SetUpsert(true)
	_,err:=database.MgoCollection.UpdateOne(context.TODO(),updateUserId,updateData,opt)
	if err!=nil{
		log.Println(err)
		return User{}
	}
	return user
}	

func MgoDeleteUser(id string) bool{
	userId,_:=strconv.Atoi(id)
	_,err:=database.MgoCollection.DeleteOne(context.TODO(),bson.D{{"id",userId}})
	if err!=nil{
		log.Println(err)
		return false
	}
	return true
}

