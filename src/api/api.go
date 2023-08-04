package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kyrinkie/admgo/src/api/middleware"
)

func Run() {
	r := gin.New()
	middleware.InitMiddleware(r)
	r.GET("/", func(c *gin.Context) {
		fmt.Println("i am in handle")
		defer fmt.Println("i am in handle defer")
		c.SetCookie("csrf_token", "dawdawd", 10, "/", "", false, true)
		//c.AbortWithStatusJSON(http.StatusTooManyRequests, "Too many requets")
		c.HTML(200, "index.html", gin.H{})
	})
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("i am in handle")
		defer fmt.Println("i am in handle defer")
		c.SetCookie("csrf_token", "dawdawd", 10, "/", "", false, true)
		//c.AbortWithStatusJSON(http.StatusTooManyRequests, "Too many requets")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ping2", func(c *gin.Context) {
		fmt.Println("i am in handle")
		defer fmt.Println("i am in handle defer")
		c.SetCookie("csrf_token", "2222", 10, "/", "", false, true)
		//c.AbortWithStatusJSON(http.StatusTooManyRequests, "Too many requets")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
