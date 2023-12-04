package main

import (
	"golangAPI/database"
	"golangAPI/middlewares"
	. "golangAPI/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
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

	engine.Use(gin.BasicAuth(gin.Accounts{"tom":"12345",}),middlewares.Logger())


	v1:=engine.Group("/v1")
	AddUserRouter(v1)

	go func(){
		database.DD()
	}()
	engine.Run(":8000")
}