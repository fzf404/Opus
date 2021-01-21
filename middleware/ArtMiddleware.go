package middleware

import (
	"Opus/model"
	"Opus/response"

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

		if len(title) < 3 || len(title) > 23 {
			response.Warning(ctx, nil, "标题字数必须在3-23个之间")
			ctx.Abort()
			return
		}
		if len(subTitle) > 120 {
			response.Warning(ctx, nil, "描述必须在120位以内")
			ctx.Abort()
			return
		}
		if len(artType) > 10 {
			response.Warning(ctx, nil, "文章类型必须在10位以内")
			ctx.Abort()
			return
		}
		if len(content) < 40 {
			response.Warning(ctx, nil, "文章字数必须大于36")
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
