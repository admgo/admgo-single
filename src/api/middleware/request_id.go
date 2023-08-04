package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func RequetID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		requestId := strings.ToUpper(strings.Replace(uuid.New().String(), "-", "", -1))
		c.Header("X-Request-Id", requestId)
		c.Set("X-Request-Id", requestId)
		fmt.Println("request")
		c.Next()
		fmt.Println("req for next")
	}
}
