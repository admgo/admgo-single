package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogToFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer fmt.Println(ctx.GetHeader("X-Request-Id"))
		fmt.Println("log")
		ctx.Next()
		fmt.Println(ctx.GetHeader("next in log"))
	}
}
