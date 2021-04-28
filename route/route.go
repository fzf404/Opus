package route

import (
	"Opus/controller"
	"Opus/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由汇总
func CollectRoute(r *gin.Engine) *gin.Engine {
	// 跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 登陆注册信息
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	artRoutes := r.Group("")
	artRoutes.Use(middleware.AuthMiddleware())
	// 获得用户信息
	artRoutes.POST("/myinfo", controller.MyInfo)
	// 添加文章
	artRoutes.POST("/addart", controller.AddArticle)
	// 修改文章
	artRoutes.POST("/modart", controller.ModArticle)
	// 删除文章
	artRoutes.POST("/delart", controller.DelArticle)
	// 上传图片
	artRoutes.POST("/upimg", controller.UploadImg)

	// 获取文章
	r.POST("/getart", controller.GetArticle)
	// 搜索
	r.POST("/search", controller.Search)
	// 获取用户信息+文章
	r.POST("/getarts", controller.GetArts)
	// 通过文章名获取文章
	r.POST("/findarts", controller.FindArticles)
	// 获取最近10条记录
	r.GET("/active", controller.GetActive)

	return r
}
