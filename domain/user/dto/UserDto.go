package dto

type UserResponse struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type UserResponseDto struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
