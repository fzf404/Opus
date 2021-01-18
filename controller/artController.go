package controller

import (
	"Opus/database"
	"Opus/dto"
	"Opus/model"
	"Opus/response"

	"github.com/gin-gonic/gin"
)

// type Article struct {
// 	gorm.Model
// 	UserID   	uint   	// 用户id
// 	Title    	string 	// 文章标题
// 	SubTitle 	string 	// 文章简介
// 	ArtType    	string 	// 文章类型
// 	Content		string 	// 文章内容
// 	Likes    	uint   	// 点赞数量
// }

// AddArticle 添加文章
func AddArticle(ctx *gin.Context) {
	DB := database.GetDB()

	user, _ := ctx.Get("user")

	userID := user.(model.User).ID
	title := ctx.PostForm("title")
	subTitle := ctx.PostForm("subtitle")
	artType := ctx.PostForm("type")
	context := ctx.PostForm("context")

	if len(title) < 2 || len(title) > 23 {
		response.Warning(ctx, nil, "标题字数必须在2-23个之间")
	}
	if len(subTitle) > 50 {
		response.Warning(ctx, nil, "描述必须在50位以内")
	}
	if len(artType) > 10 {
		response.Warning(ctx, nil, "文章类型必须在10位以内")
	}
	if len(context) < 20 {
		response.Warning(ctx, nil, "文章字数必须大于18")
	}

	newArt := model.Article{
		UserID:   userID,
		Title:    title,
		SubTitle: subTitle,
		ArtType:  artType,
		Content:  context,
		Likes:    0,
	}

	DB.Create(&newArt)

	response.Success(ctx, nil, "文章发布成功")

}

// GetArticle 获取文章
func GetArticle(ctx *gin.Context) {
	DB := database.GetDB()

	artID := ctx.PostForm("artid")

	var art = model.Article{}
	DB.First(&art, artID)
	if art.ID == 0 {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	response.Success(ctx, gin.H{"article": dto.ArticleDto(art)}, "获取成功")
}
