package Inputs

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
