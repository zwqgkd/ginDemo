package main

import (
	"golangAPI/database"
	"golangAPI/middlewares"
	."golangAPI/src"
	"io"
	"os"
	"golangAPI/pojo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLogging(){
	f,_:=os.Create("gin.log")
	//gin的日志写到文件和终端
	gin.DefaultWriter=io.MultiWriter(f,os.Stdout)
}

func main(){
	setupLogging()
	//Engine是RouterGroup的子类
	engine:=gin.Default()
	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("userpasd",middlewares.UserPasd)
		v.RegisterStructValidation(middlewares.UserList,pojo.Users{})
	}

	engine.Use( 
		//gin.BasicAuth(gin.Accounts{"tom":"12345",}),
		middlewares.Logger(),//log格式
		gin.Recovery(),
		middlewares.SetSession(),
	)


	v1:=engine.Group("/v1")
	AddUserRouter(v1)

	go func(){
		database.DD()
		database.MD()
	}()
	engine.Run(":8000")
}