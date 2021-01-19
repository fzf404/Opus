package controller

import (
	"Opus/database"
	"Opus/dto"
	"Opus/model"
	"Opus/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddArticle 添加文章
func AddArticle(ctx *gin.Context) {
	DB := database.GetDB()

	user, _ := ctx.Get("user")

	userID := user.(model.User).ID
	title := ctx.PostForm("title")
	subTitle := ctx.PostForm("subtitle")
	artType := ctx.PostForm("type")
	headImg := ctx.PostForm("headimg")
	context := ctx.PostForm("context")

	if len(title) < 2 || len(title) > 23 {
		response.Warning(ctx, nil, "标题字数必须在2-23个之间")
		return
	}
	if len(subTitle) > 50 {
		response.Warning(ctx, nil, "描述必须在50位以内")
		return
	}
	if len(artType) > 10 {
		response.Warning(ctx, nil, "文章类型必须在10位以内")
		return
	}
	if len(context) < 40 {
		response.Warning(ctx, nil, "文章字数必须大于36")
		return
	}

	newArt := model.Article{
		UserID:   userID,
		Title:    title,
		SubTitle: subTitle,
		ArtType:  artType,
		HeadImg:  headImg,
		Content:  context,
		Likes:    0,
	}

	DB.Create(&newArt)

	response.Success(ctx, nil, "文章发布成功")

}

// GetArticle 获取文章
func GetArticle(ctx *gin.Context) {
	DB := database.GetDB()

	// 获取文章
	artID := ctx.PostForm("artid")
	var art model.Article
	DB.First(&art, artID)
	if art.ID == 0 {
		response.NotFind(ctx, nil, "文章不存在")
		return
	}
	response.Success(ctx, gin.H{"article": dto.ArticleDto(art)}, "获取成功")
}

// GetInfo 取全部文章列表
func GetInfo(ctx *gin.Context) {

	DB := database.GetDB()

	// 用户信息处理
	var user model.User
	userid := ctx.PostForm("userid")
	DB.Where("id = ?", userid).First(&user)
	if user.ID == 0 {
		response.NotFind(ctx, nil, "该用户不存在")
		return
	}

	// map处理全部articles
	var articles []model.Article
	items := make(map[string]model.ArticleDto)
	// 获取列表
	DB.Where("user_id = ?", userid).Find(&articles)
	log.Print(articles)
	// 判断获取情况
	if len(articles) == 0 {
		response.NotFind(ctx, nil, "未找到已发布的文章")
		return
	}

	for index, art := range articles {
		// 使用dto处理全部article
		items["article"+strconv.Itoa(index)] = dto.ArticleInfoDto(art)
	}
	response.Success(ctx, gin.H{
		"user":     dto.TouserUserDto(user),
		"articles": items,
	}, "获取全部文章成功")
}
