package controller

import (
	"Opus/model"
	"Opus/response"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// UploadImg 上传图片
func UploadImg(ctx *gin.Context) {
	// 图片验证
	file, err := ctx.FormFile("img")
	if err != nil {
		response.Fail(ctx, nil, "图片获取失败")
		return
	}
	user, _ := ctx.Get("user")
	path := viper.GetString("common.path")
	now := time.Now().Unix()
	file.Filename = fmt.Sprintf("%s%s%d%s", path, user.(model.User).Name, now, file.Filename)
	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		response.Fail(ctx, nil, "图片上传失败")
		return
	}
	response.Success(ctx, gin.H{"url": file.Filename[4:]}, "上传成功")
}
