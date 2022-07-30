package Inputs

type RegisterUserInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Occupation      string `json:"occupation" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword" binding:"eqfield=Password,required"`
}

type LoginUserInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmail struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}
