package route

import (
	"Opus/controller"
	"Opus/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由汇总
func CollectRoute(r *gin.Engine) *gin.Engine {
	
	// 登陆注册
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 中间件验证权限
	r.POST("/myinfo", middleware.AuthMiddleware(), controller.MyInfo)
	// 获取用户信息
	r.POST("/getinfo", controller.GetInfo)

	// 添加文章
	r.POST("/addart", middleware.AuthMiddleware(), controller.AddArticle)
	// 获取文章
	r.POST("/getart", controller.GetArticle)
	// 获取全部文章
	r.POST("/allart", controller.GetAllArt)
	
	return r
}
