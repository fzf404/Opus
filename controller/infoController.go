package controller

import (
	"Opus/database"
	"Opus/dto"
	"Opus/model"
	"Opus/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// GetArticle 获取文章
func GetArticle(ctx *gin.Context) {
	db := database.GetDB()

	// 获取文章
	artID := ctx.PostForm("artid")
	var art model.Article
	db.First(&art, artID)
	if art.ID == 0 {
		response.NotFind(ctx, nil, "文章不存在")
		return
	}
	// 获取用户
	var user model.User
	db.First(&user, art.UserID)
	if user.ID == 0 {
		response.NotFind(ctx, nil, "用户不存在")
		return
	}
	response.Success(ctx, gin.H{"user": dto.TouserUserDto(user), "article": dto.ArticleDto(art)}, "获取成功")
}

// GetArts 取全部文章列表
func GetArts(ctx *gin.Context) {

	db := database.GetDB()

	// 用户信息处理
	var user model.User
	userid := ctx.PostForm("userid")
	db.Where("id = ?", userid).First(&user)
	if user.ID == 0 {
		response.NotFind(ctx, nil, "该用户不存在")
		return
	}

	// map处理全部articles
	var articles []model.Article
	items := make(map[string]model.ArticleDto)

	getmaxart := viper.GetString("common.maxart")
	maxart, _ := strconv.Atoi(getmaxart)
	// 获取列表
	db.Where("user_id = ?", userid).Order("id desc").Limit(maxart).Find(&articles)
	// 判断获取情况
	if len(articles) == 0 {
		response.NotFind(ctx, nil, "未找到已发布的文章")
		return
	}

	for index, art := range articles {
		// 使用dto处理全部article
		items["article"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}
	response.Success(ctx, gin.H{
		"user":     dto.TouserUserDto(user),
		"articles": items,
	}, "获取全部文章成功")
}

// Search 搜索用户名或文章名
func Search(ctx *gin.Context) {
	db := database.GetDB()

	name := ctx.PostForm("name")
	if len(name) < 3 {
		response.NotFind(ctx, nil, "什么都没找到~")
		return
	}

	var user model.User
	var article model.Article
	db.First(&user, "name = ?", name)
	db.First(&article, "title LIKE ?", "%"+name+"%")

	if user.ID != 0 {
		response.Success(ctx, gin.H{
			"userid": user.ID,
		}, "找到用户了")
		return
	}
	if article.ID != 0 {
		response.Success(ctx, gin.H{
			"userid": article.UserID,
		}, "找到文章了")
		return
	}
	response.NotFind(ctx, nil, "什么都没找到~")
}

// FindArticles 通过关键词查找文章
func FindArticles(ctx *gin.Context) {

	db := database.GetDB()
	name := ctx.PostForm("name")
	// map处理全部articles
	var articles []model.Article
	items := make(map[string]model.ArticleDto)
	// 获取列表
	db.Where("title LIKE ?", "%"+name+"%").Find(&articles)
	// 判断获取情况
	if len(articles) == 0 {
		response.NotFind(ctx, nil, "未找到已发布的文章")
		return
	}

	for index, art := range articles {
		// 使用dto处理全部article
		items["article"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}

	response.Success(ctx, gin.H{
		"articles": items,
	}, "获取全部文章成功")
}

// GetActive 获取最近的文章
func GetActive(ctx *gin.Context) {
	db := database.GetDB()
	// 获取参数
	datacode := ctx.Query("datacode")
	getmaxart := viper.GetString("common.maxart")
	maxart, _ := strconv.Atoi(getmaxart)

	var articles []model.Article
	items := make(map[string]model.ArticleDto)

	switch datacode {
	case "0":
		db.Order("id desc").Limit(maxart).Find(&articles)
	case "1":
		db.Where("super = ?", 1).Order("id desc").Limit(maxart).Find(&articles)
	case "2":
		db.Where("art_type = ?", "科技").Order("id desc").Limit(maxart).Find(&articles)
	case "3":
		db.Where("art_type = ?", "校园").Order("id desc").Limit(maxart).Find(&articles)
	case "4":
		db.Where("art_type = ?", "艺术").Order("id desc").Limit(maxart).Find(&articles)
	case "10":
		db.Where("art_type = ?", "未分类").Order("id desc").Limit(maxart).Find(&articles)
	default:
		db.Order("id desc").Limit(maxart).Find(&articles)
	}

	if len(articles) == 0 {
		response.NotFind(ctx, nil, "未找到已发布的文章")
		return
	}

	for index, art := range articles {
		// 使用dto处理全部article
		items["article"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}

	// 获取最近的文章
	response.Success(ctx, gin.H{
		"articles": items,
	}, "获取全部文章成功")
}
