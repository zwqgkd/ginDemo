package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userKey = "session_id"

// Use cookie to store session id
func SetSession() gin.HandlerFunc{
	store:=cookie.NewStore([]byte("secret"))
	return sessions.Sessions("mysession",store)
}


//User Auth Session Middleware
func AuthSession() gin.HandlerFunc{
	return func(c *gin.Context){
		session:=sessions.Default(c)
		sessionID:=session.Get(userKey)
		if sessionID==nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"message":"Unauthorized",
			})
			return
		}
		c.Next()
	}
}

//Save Session For User
func SaveSession(c *gin.Context, userID int){
	session:=sessions.Default(c)
	session.Set(userKey,userID)
	session.Save()
}

//Clear Session for User
func ClearSession(c *gin.Context){
	session:=sessions.Default(c)
	session.Clear()
	session.Save()
}

//Get Session for User
func GetSession(c *gin.Context)int {
	session:=sessions.Default(c)
	sessionID:=session.Get(userKey)
	if sessionID==nil{
		return 0
	}
	return sessionID.(int)
}

//Check Session for User
func CheckSession(c *gin.Context) bool{
	session:=sessions.Default(c)
	sessionID:=session.Get(userKey)
	return sessionID!=nil
}