package model

import (
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	HeadImg  string
	Super    bool
	Active   bool
}

// UserDto 用户数据传输
type UserDto struct {
	UserID  uint   `json:"userid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	HeadImg string `json:"headimg"`
	Super   bool   `json:"Super"`
	Active  bool   `json:"active"`
}
