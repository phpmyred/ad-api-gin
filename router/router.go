package router

import (
	"github.com/WangaduoApi/ad-api-gin/router/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter
// @Description  主路由
// @Author aDuo 2024-08-14 22:27:31
// @Return *gin.Engine
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 加入设置请求头中间件  跟 错误捕获中间件
	r.Use(middleware.CORSMiddleware(), middleware.RecoverMiddleware())
	api := r.Group("/api")

	//r.GET("/ping", api.Index)

	// 加载APP的路由
	appRouter(api.Group("/app"))

	adminRouter(api.Group("/admin"))

	return r
}
