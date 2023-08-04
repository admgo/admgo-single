package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer fmt.Println("i am in auth defer")
		fmt.Println("auth")
		// ctx.Abort()
	}
}
