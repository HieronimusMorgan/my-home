package utils

import (
	"Master_Data/module/domain"
	"Master_Data/module/dto/out"
)

func ConvertUserToResponse(user domain.User) out.UserResponse {
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
