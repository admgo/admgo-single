package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(e *gin.Engine) {
	e.Use(LogToFile()) // 写入日志文件
	e.Use(RequetID())
	e.Use(Auth()) // 写入日志文件
}
