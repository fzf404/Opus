package route

import (
	"Opus/controller"
	"Opus/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由汇总
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	// 中间件验证权限
	r.POST("/info", middleware.AuthMiddleware(),controller.Info)
	// 添加文章
	r.POST("/addart",middleware.AuthMiddleware(),controller.AddArticle)
	// 获取文章
	r.POST("/getart",controller.GetArticle)
	return r
}
