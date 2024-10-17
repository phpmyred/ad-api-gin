package middleware

import (
	"fmt"
	"github.com/WangaduoApi/ad-api-gin/api/contextData/response"
	"github.com/gin-gonic/gin"
)

// RecoverMiddleware 错误捕获中间件
// @Description  接口只需要任意用 panic 该方法会自动捕获
// @Author aDuo 2024-08-14 22:53:25
// @Return gin.HandlerFunc

func RecoverMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if fmt.Sprint(err) != "" {
					response.Fail(ctx, fmt.Sprint(err), nil)
				} else {
					response.Error(ctx)
				}
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
