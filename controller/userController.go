package controller

import (
	"Opus/database"
	"Opus/dto"
	"Opus/middleware"
	"Opus/model"
	"Opus/response"
	"Opus/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register 用户注册信息处理
func Register(ctx *gin.Context) {

	DB := database.GetDB()

	// 获得用户数据
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	// 信息验证
	if len(name) < 3 || len(name) > 12 {
		response.Warning(ctx, nil, "用户名须在3-12位之间")
		return
	}
	if len(password) < 6 || len(password) > 16 {
		response.Warning(ctx, nil, "密码须在6-16位之间")
		return
	}

	// 判断邮箱是否合法
	if !util.VerifyEmailFormat(email) {
		response.Fail(ctx, nil, "邮箱不合法")
		return
	}

	// 判断用户名邮箱是否存在
	var user model.User
	DB.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		response.Warning(ctx, nil, "用户名已存在")
		return
	}
	DB.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		response.Fail(ctx, nil, "邮箱已存在")
		return
	}

	// 创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}
	newUser := model.User{
		Name:     name,
		Email:    email,
		Password: string(hashPassword),
		HeadImg:  "https://i0.hdslb.com/bfs/face/member/noface.jpg",
		Super:    false,
		Active:   false,
	}
	DB.Create(&newUser)

	response.Success(ctx, nil, "注册成功")
}

// Login 用户登录信息处理
func Login(ctx *gin.Context) {
	DB := database.GetDB()

	//获取数据
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	// 判断密码
	if len(password) < 6 || len(password) > 16 {
		response.Fail(ctx, nil, "密码错误")
		return
	}

	// 判断用户是否存在
	var user model.User
	DB.Where("name = ?", name).First(&user)
	// 用户名判断
	if user.ID == 0 {
		// 邮箱判断
		DB.Where("email = ?", name).First(&user)
		if user.ID == 0 {
			response.Fail(ctx, nil, "用户名或密码错误")
			return
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "用户名或密码错误")
		return
	}
	// 分发Token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "Token发放失败")
		log.Print("token generate error:", err)
		return
	}

	response.Success(ctx, gin.H{"token": token}, "登陆成功")

}

// MyInfo 输出我的信息
func MyInfo(ctx *gin.Context) {
	// 获取中间件添加的user
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.TouserMyDto(user.(model.User))}, "个人信息获取成功")

}

// Search 搜索用户名或文章名
func Search(ctx *gin.Context) {
	DB := database.GetDB()

	name := ctx.PostForm("name")
	if len(name) < 3 {
		response.NotFind(ctx, nil, "什么都没找到~")
		return
	}

	var user model.User
	var article model.Article
	DB.First(&user, "name = ?", name)
	DB.First(&article, "title = ?", name)

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
