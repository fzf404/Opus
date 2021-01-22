package dto

import (
	"Opus/model"
)

// ArticleDto 文章Dto
func ArticleDto(art model.Article) model.ArticleDto {
	timeTemplate := "2006-01-02 15:04:05"
	// creatTime, _ := time.Parse(timeTemplate, art.CreatedAt.Format(timeTemplate))
	return model.ArticleDto{
		CreateAt: art.CreatedAt.Local().Format(timeTemplate),
		UpdateAt: art.UpdatedAt.Local().Format(timeTemplate),
		ArtID:    art.ID,
		UserID:   art.UserID,
		UserName: art.UserName,
		Title:    art.Title,
		SubTitle: art.SubTitle,
		HeadImg:  art.HeadImg,
		ArtType:  art.ArtType,
		Content:  art.Content,
		Share:    art.Share,
		Likes:    art.Likes,
		Super:    art.Super,
	}
}

// ArticleInfoDto 文章Dto
func ArticleInfoDto(art model.Article) model.ArticleDto {
	// 除Context
	return model.ArticleDto{
		ArtID:    art.ID,
		UserID:   art.UserID,
		UserName: art.UserName,
		Title:    art.Title,
		SubTitle: art.SubTitle,
		HeadImg:  art.HeadImg,
		ArtType:  art.ArtType,
		Share:    art.Share,
		Likes:    art.Likes,
		Super:    art.Super,
	}
}
