package utils

import (
	"Master_Data/module/domain/master"
	"Master_Data/module/dto/out"
)

func ConvertUserToResponse(user master.User) out.UserResponse {
	return out.UserResponse{
		ID:             user.UserID,
		ClientID:       user.ClientID,
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		ProfilePicture: user.ProfilePicture,
	}
}
