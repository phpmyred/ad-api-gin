package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// DecryptMiddleware 请求参数解密中间件 解密只针对 POST 方法
// @Description 会将请求的参数解密后塞回 BODY当中
// @Author aDuo 2024-08-14 22:12:29 ${time}
// @Return gin.HandlerFunc
func DecryptMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" {
			if viper.GetBool("encipher.IsEncrypt") {
				ctx.Set("encryption", "yes")
			}
		}

		ctx.Next()
	}
}
