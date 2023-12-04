package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


//Get User
func FindAllUsers(c *gin.Context){
	// c.JSON(http.StatusOK,userList)
	users:=pojo.FindAllUsers()
	c.JSON(http.StatusOK,users)
}

//Get User by id
func FindByUserId(c *gin.Context){
	user:=pojo.FindByUserId(c.Param("id"))
	if user.Id==0{
		c.JSON(http.StatusNotFound,"User not found")
		return 
	}
	log.Println("user->",user)
	c.JSON(http.StatusOK,user)
}

//Post User
func PostUser(c *gin.Context){
	user:=pojo.User{}
	err:=c.BindJSON(&user)
	if err!=nil{
		c.JSON(http.StatusNotAcceptable,"Error:"+err.Error())
		return
	}
	//userList=append(userList,user)
	newUser:=pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

//delete User
func DeleteUser(c *gin.Context){
	//userId:=c.Param("id")//注意返回的类型是string
	flag:=pojo.DeleteUser(c.Param("id"))
	if !flag{
		c.JSON(http.StatusNotFound,"User not found")
		return 
	}
	c.JSON(http.StatusOK,"success")
}

//put User
func Putuser(c *gin.Context){
	user:=pojo.User{}
	err:=c.BindJSON(&user)
	if err!=nil{
		c.JSON(http.StatusNotAcceptable,"Error:"+err.Error())
		return
	}
	user=pojo.UpdateUser(c.Param("id"),user)
	if user.Id==0{
		c.JSON(http.StatusNotFound,"User not found")
		return 
	}
	c.JSON(http.StatusOK,user)
}