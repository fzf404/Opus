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
	Super    bool
	Active   bool
}

// UserDto 用户数据传输
type UserDto struct {
	UserID uint   `json:"userID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Super  bool   `json:"bool"`
	Active bool   `json:"active"`
}

