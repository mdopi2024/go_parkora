package users

import (
	userdto "parkora/internal/users/dto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Role     string `gorm:"type:varchar(20);default:driver;not null" json:"role"`
}

func (u *User) ToResponse() userdto.UserResponse {
	return userdto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToCreateResponse(message string) userdto.RegisterUserResponse {
	return userdto.RegisterUserResponse{
		Success: true,
		Message: message,
		Data:    u.ToResponse(),
	}
}

func (u *User) ToLoginResponse(token, message string) userdto.LoginUserResponse {
	return userdto.LoginUserResponse{
		Success: true,
		Message: message,
		Data: userdto.LoginUserData{
			Token: token,
			User:  u.ToResponse(),
		},
	}
}
