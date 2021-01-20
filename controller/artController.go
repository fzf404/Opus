package controller

import (
	"Opus/database"
	"Opus/model"
	"Opus/response"

	"github.com/gin-gonic/gin"
)

// AddArticle 添加文章
func AddArticle(ctx *gin.Context) {
	DB := database.GetDB()

	// 获得文章信息及用户信息
	user, _ := ctx.Get("user")
	article, _ := ctx.Get("article")

	// 新建文章
	newArt := model.Article{
		UserID:   user.(model.User).ID,
		UserName: user.(model.User).Name,
		Title:    article.(model.Article).Title,
		SubTitle: article.(model.Article).SubTitle,
		ArtType:  article.(model.Article).ArtType,
		HeadImg:  article.(model.Article).HeadImg,
		Content:  article.(model.Article).Content,
		Likes:    0,
	}

	DB.Create(&newArt)

	response.Success(ctx, gin.H{"name": user.(model.User).Name, "title": newArt.Title, "artid": newArt.ID}, "文章发布成功")

}

// DelArticle 删除文章
func DelArticle(ctx *gin.Context) {
	DB := database.GetDB()

	user, _ := ctx.Get("user")

	// 获取用户即文章信息
	userID := user.(model.User).ID
	artID := ctx.PostForm("artid")
	if len(artID) == 0 {
		response.Fail(ctx, nil, "删除请求非法")
		return
	}

	var article model.Article
	DB.First(&article, artID)
	if article.UserID != userID {
		response.Fail(ctx, nil, "删除请求非法")
		return
	}

	DB.Delete(&article)
	response.Success(ctx, gin.H{"name": user.(model.User).Name, "title": article.Title, "artid": article.ID}, "删除成功")
}

// ModArticle 修改文章
func ModArticle(ctx *gin.Context) {
	DB := database.GetDB()

	user, _ := ctx.Get("user")
	// 获取用户id和文章id
	userID := user.(model.User).ID
	artID := ctx.PostForm("artid")

	// 判断该文章是否为该用户所有
	var art model.Article
	DB.First(&art, artID)
	if art.UserID != userID {
		response.Fail(ctx, nil, "修改请求非法")
		return
	}

	article, _ := ctx.Get("article")

	art.Title = article.(model.Article).Title
	art.SubTitle = article.(model.Article).SubTitle
	art.ArtType = article.(model.Article).ArtType
	art.HeadImg = article.(model.Article).HeadImg
	art.Content = article.(model.Article).Content

	DB.Save(&art)
	response.Success(ctx, gin.H{"name": user.(model.User).Name, "title": art.Title, "artid": art.ID}, "文章更新成功")
}
