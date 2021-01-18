package dto

import "Opus/model"

// ArticleDto 文章Dto
func ArticleDto(art model.Article) model.ArticleDto {

	return model.ArticleDto{
		ArtID:    art.ID,
		UserID:   art.UserID,
		Title:    art.Title,
		SubTitle: art.SubTitle,
		ArtType:  art.ArtType,
		Content:  art.Content,
		Likes:    art.Likes,
	}
}
// ArticleDto 文章Dto
func ArticleInfoDto(art model.Article) model.ArticleDto {

	return model.ArticleDto{
		ArtID:    art.ID,
		UserID:   art.UserID,
		Title:    art.Title,
		SubTitle: art.SubTitle,
		ArtType:  art.ArtType,
		Likes:    art.Likes,
	}
}
