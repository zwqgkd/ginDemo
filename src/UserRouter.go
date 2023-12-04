package src

import (
	"golangAPI/service"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup){
	user:=r.Group("/users")
	user.GET("/:id",service.FindByUserId)
	user.GET("/",service.FindAllUsers)
	user.POST("/",service.PostUser)
	user.DELETE("/:id",service.DeleteUser)
	user.PUT("/:id",service.Putuser)
}