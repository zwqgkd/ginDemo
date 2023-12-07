package src

import (
	session "golangAPI/middlewares"
	"golangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	//users have common middlewares session.SetSession()
	user := r.Group("/users", session.SetSession())

	user.GET("/:id", service.FindByUserId)
	user.GET("/", service.FindAllUsers)
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)

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
