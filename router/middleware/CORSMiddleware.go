package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 设置请求头
// @Description  主要作用是允许跨域 等设置
// @Author aDuo 2024-08-14 22:52:50
// @Return gin.HandlerFunc

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置 Access-Control-Allow-Origin 头部，允许来自任何源的请求。
		// 注意：使用 "*" 允许任何源访问。
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置 Access-Control-Max-Age 头部，指定预检请求的结果可以被缓存的时间（秒）。
		// 这里设置为 86400 秒（24 小时）。
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// 设置 Access-Control-Allow-Methods 头部，指定允许的 HTTP 方法。
		// "*" 表示允许所有方法。
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		// 设置 Access-Control-Allow-Headers 头部，指定允许的请求头。
		// "*" 表示允许所有请求头。
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		// 设置 Access-Control-Allow-Credentials 头部，指示是否允许在跨源请求中使用凭据（如 cookies）。
		// 设置为 "true" 允许携带凭据。
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果请求方法是 OPTIONS，则表示这是一个预检请求。
		// 响应 200 OK 状态码并中止进一步处理请求。
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
