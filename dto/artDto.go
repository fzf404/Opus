package dto

import "Opus/model"

// ArticleDto 文章Dto
func ArticleDto(art model.Article) model.ArticleDto {

	return model.ArticleDto{
		UserID:   art.UserID,
		Title:    art.Title,
		SubTitle: art.SubTitle,
		ArtType:  art.ArtType,
		Content:  art.Content,
		Likes:    art.Likes,
	}
}
