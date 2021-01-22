package model

import "gorm.io/gorm"

// Article 文章数据
type Article struct {
	gorm.Model
	UserID   uint   // 用户id
	UserName string // 用户名
	Title    string // 文章标题
	SubTitle string // 文章简介
	ArtType  string // 文章类型
	HeadImg  string // 首页图
	Content  string // 文章内容
	Likes    uint   // 点赞数量
	Share    uint   // 分享数量
	Super    bool   // 加精
}

// ArticleDto 文章数据传输
type ArticleDto struct {
	CreateAt string `json:"create"`   // 创建时间
	UpdateAt string `json:"update"`   // 修改事件
	ArtID    uint   `json:"artid"`    // 文章id
	UserID   uint   `json:"userid"`   // 用户id
	UserName string `json:"username"` // 用户名
	Title    string `json:"title"`    // 文章标题
	SubTitle string `json:"subtitle"` // 文章简介
	ArtType  string `json:"arttype"`  // 文章类型
	HeadImg  string `json:"headimg"`  // 首页图
	Content  string `json:"content"`  // 文章内容
	Likes    uint   `json:"likes"`    // 点赞数量
	Share    uint   `json:"share"`
	Super    bool   `json:"super"`
}
