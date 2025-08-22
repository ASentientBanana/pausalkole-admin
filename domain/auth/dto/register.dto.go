package dto

type RegisterDTO struct {
	Email           string `json:"email"`
	Password        string `json:"password" binding:"required" validate:"min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"min=8"`
}
