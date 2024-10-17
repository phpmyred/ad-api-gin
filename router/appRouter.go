package router

import (
	"github.com/WangaduoApi/ad-api-gin/api/app/captcha"
	"github.com/gin-gonic/gin"
)

// appRouter
// @Description APP的路由组
// @Author aDuo 2024-08-15 00:56:39
// @Param r
func appRouter(r *gin.RouterGroup) {

	r.POST("/captcha/get", captcha.Get)
	r.POST("/captcha/check", captcha.Check)
}
