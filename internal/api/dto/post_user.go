package dto

type PostUser struct {
	Email    string `json:"email" validate:"email"`
	Login    string `json:"login"`
	Password string `json:"password" validate:"required"`
}
