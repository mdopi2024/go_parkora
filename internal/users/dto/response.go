package dto

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUserResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

type LoginUserResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    LoginUserData `json:"data"`
}

type LoginUserData struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
