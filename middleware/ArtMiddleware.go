package middleware

import (
	"Opus/model"
	"Opus/response"
	"log"

	"github.com/gin-gonic/gin"
)

// ArtMiddleware 判断文章是否合法
func ArtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		subTitle := ctx.PostForm("subtitle")
		artType := ctx.PostForm("type")
		// headImg := ctx.PostForm("headimg")
		headImg := "/static/img/noimg.jpg"
		content := ctx.PostForm("content")

		if len(title) < 6 || len(title) > 64 {
			response.Warning(ctx, nil, "标题字数必须在3-23个之间")
			ctx.Abort()
			return
		}
		if len(subTitle) > 256 {
			response.Warning(ctx, nil, "描述必须在100位以内")
			ctx.Abort()
			return
		}
		if len(artType) > 16 {
			log.Print(artType)
			log.Print(len(artType))
			response.Warning(ctx, nil, "文章类型必须在5位以内")
			ctx.Abort()
			return
		}
		if len(content) < 512 {
			response.Warning(ctx, nil, "文章字数必须大于100")
			ctx.Abort()
			return
		}

		article := model.Article{
			Title:    title,
			SubTitle: subTitle,
			ArtType:  artType,
			HeadImg:  headImg,
			Content:  content,
		}
		ctx.Set("article", article)
		ctx.Next()
	}
}
