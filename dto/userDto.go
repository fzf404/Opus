package dto

import "Opus/model"

// TouserDto 用户Dto
func TouserDto(user model.User) model.UserDto {

	return model.UserDto{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	}

}
