package src

import (
	session "golangAPI/middlewares"
	"golangAPI/service"
	"golangAPI/pojo"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	//users have common middlewares session.SetSession()
	user := r.Group("/users", session.SetSession())
	//user.GET("/",service.FindAllUsers)
	user.GET("/",service.CachUserAllDecorator(service.RedisAllUser,"user_all",[]pojo.User{}))
	//user.GET("/:id", service.FindByUserId)
	user.GET("/:id",service.CachOneUserDecorator(service.RedisOneUser, "id", "user_%s",pojo.User{}))
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)

	//MongoDB ----------------------------------------------------
	mgo:=user.Group("/mongo")
	mgo.GET("/", service.MgoDBFindAllUser)
	mgo.GET("/:id", service.MgoDBFindOneUser)
	mgo.PUT("/:id",service.MgoDBPutUser)
	mgo.POST("/", service.MgoDBCreateUser)
	mgo.DELETE("/:id", service.MgoDBDeleteUser)

	//put user
	user.PUT("/:id", service.Putuser)
	//Login
	user.POST("/login", service.LoginUser)

	//Check User Session
	user.GET("/check", service.CheckUserSession)

	//此后的都加上了AhthSession中间件
	user.Use(session.AuthSession())
		//Delete
	user.DELETE("/:id", service.DeleteUser)
	user.GET("/logout", service.LogoutUser)


}
