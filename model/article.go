package model

import "gorm.io/gorm"

// Article 文章数据
type Article struct {
	gorm.Model
	UserID   uint   // 用户id
	Title    string // 文章标题
	SubTitle string // 文章简介
	ArtType  string // 文章类型
	Content  string // 文章内容
	Likes    uint   // 点赞数量
}

// ArticleDto 文章数据传输
type ArticleDto struct {
	UserID   uint   `json:"userid"`   // 用户id
	Title    string `json:"title"`    // 文章标题
	SubTitle string `json:"subtitle"` // 文章简介
	ArtType  string `json:"arttype"`  // 文章类型
	Content  string `json:"content"`  // 文章内容
	Likes    uint   `json:"likes"`    // 点赞数量
}
