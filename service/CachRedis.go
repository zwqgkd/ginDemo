package service

import (
	"fmt"
	red "golangAPI/database"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
)

func CachOneUserDecorator(h gin.HandlerFunc, porm string , readKeyPattern string, empty interface{}) gin.HandlerFunc{
	return func(c *gin.Context){
		keyId:=c.Param(porm)
		redisKey:=fmt.Sprintf(readKeyPattern,keyId)
		conn:=red.RedisDefaultPool.Get()
		defer conn.Close()
		data,err:=redis.Bytes(conn.Do("GET",redisKey))
		if err!=nil{
			h(c)
			dbResult,exist:=c.Get("dbResult")
			if !exist{
				dbResult=empty
			}
			redisData,_:=ffjson.Marshal(dbResult)
			conn.Do("SETEX",redisKey,30,redisData)
			c.JSON(http.StatusOK,gin.H{
				"message":"From DB",
				"data":dbResult,
			})
			return
		}
		ffjson.Unmarshal(data,&empty)
		c.JSON(http.StatusOK,gin.H{
			"message":"From Redis",
			"data":empty,
		})
	}
}


func CachUserAllDecorator(h gin.HandlerFunc, readKey string, empty interface{}) gin.HandlerFunc{
	return func(c *gin.Context){
		conn:=red.RedisDefaultPool.Get()
		defer conn.Close()

		data,err:=redis.Bytes(conn.Do("GET",readKey))
		if err!=nil{
			h(c)
			dbUserAll,exist:=c.Get("dbUserAll")
			if !exist{
				dbUserAll=empty
			}
			redisData,_:=ffjson.Marshal(dbUserAll)
			conn.Do("SETEX",readKey,30,redisData)
			c.JSON(http.StatusOK,gin.H{
				"message":"From DB",
				"data":dbUserAll,
			})
			return 
		}
		ffjson.Unmarshal(data,&empty)
		c.JSON(http.StatusOK,gin.H{
			"message":"From Redis",
			"data":empty,
		})
	}
}