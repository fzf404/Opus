package dto

import "Opus/model"

// TouserMyDto 登录用户Dto
func TouserMyDto(user model.User) model.UserDto {

	return model.UserDto{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		HeadImg: user.HeadImg,
	}
}

// TouserUserDto 默认Dto
func TouserUserDto(user model.User) model.UserDto {

	return model.UserDto{
		UserID: user.ID,
		Name:   user.Name,
		HeadImg: user.HeadImg,
	}

}
