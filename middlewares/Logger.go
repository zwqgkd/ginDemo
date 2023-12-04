package middlewares 

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Logger() gin.HandlerFunc{
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string{
			return fmt.Sprintf("%s - [%s] %s %s %d \n",
			params.ClientIP,
			params.TimeStamp,
			params.Method,
			params.Path,
			params.StatusCode)
		},
	)
}