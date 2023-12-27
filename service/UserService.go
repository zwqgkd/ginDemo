package service

import (
	"golangAPI/middlewares"
	"golangAPI/pojo"
	DB "golangAPI/database"
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
	log.Println("user->",user)
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


//CreateUserList
func CreateUserList(c *gin.Context){
	users:=pojo.Users{}
	err:=c.BindJSON(&users)
	if err!=nil{
		c.String(400,"Error:%s", err.Error())
		return 
	}
	c.JSON(http.StatusOK,users)
}

//Login User
func LoginUser(c *gin.Context){
	name:=c.PostForm("name")
	password:=c.PostForm("password")
	user:=pojo.CheckUserPassword(name,password)
	if user.Id==0{
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	middlewares.SaveSession(c,user.Id)
	c.JSON(http.StatusOK,gin.H{
		"message":"Login Success",
		"user":user,
		"session":middlewares.GetSession(c),
	})
}

//Logout User
func LogoutUser(c *gin.Context){
	middlewares.ClearSession(c)
	c.JSON(http.StatusOK,gin.H{
		"message":"Logout Success",
	})
}


//Check User Session
func CheckUserSession(c *gin.Context){
	sessionID:=middlewares.GetSession(c)
	if sessionID==0{
		c.JSON(http.StatusUnauthorized,"Error")
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Check Session Success",
		"sessionID":sessionID,
	})
}


//Redis User
func RedisOneUser(c *gin.Context){
	id:=c.Param("id")
	if id=="0"{
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	user:=pojo.User{}
	DB.DBConnect.Find(&user,id)
	c.Set("dbResult",user)
}

//Redis All User
func RedisAllUser(c *gin.Context){
	users:=[]pojo.User{}
	DB.DBConnect.Find(&users)
	c.Set("dbUserAll",users)
}

//mongoDB---------------------------------------------------------------
//MgoDB create user
func MgoDBCreateUser(c *gin.Context){
	user:=pojo.User{}
	err:=c.BindJSON(&user)
	if err!=nil{
		c.JSON(http.StatusNotAcceptable,"Error:"+err.Error())
		return
	}
	newUser:=pojo.MgoCreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message":"Create User Successfully",
		"user": newUser,
	})
}

func MgoDBFindAllUser(c *gin.Context){
	users:=pojo.MgoFindAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"message":"Find All User Successfully",
		"user": users,
	})
}

func MgoDBFindOneUser(c *gin.Context){
	user:=pojo.MgoFindById(c.Param("id"))
	if user.Id==0{
		c.JSON(http.StatusNotFound,"User not found")
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Find User Successfully",
		"user": user,
	})
}

//MongoDB Put User
func MgoDBPutUser(c *gin.Context){
	user:=pojo.User{}
	err:=c.BindJSON(&user)
	if err!=nil{
		c.JSON(http.StatusNotAcceptable,"Error:"+err.Error())
		return
	}
	user = pojo.MgoPutUser(c.Param("id"),user)
	if user.Id==0{
		c.JSON(http.StatusNotFound,"User not found")
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Update User Successfully",
		"user":user,
	})
}

//MongoDB Delete User 
func MgoDBDeleteUser(c *gin.Context){
	isDeleted:=pojo.MgoDeleteUser(c.Param("id"))
	if !isDeleted{
		c.JSON(http.StatusNotFound,"User not found")
		return
	}
	c.JSON(http.StatusOK,"success")
}