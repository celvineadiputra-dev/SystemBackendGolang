package Responses

import (
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Models/Users"
	"strconv"
)

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResponseRegular struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserFormat(user Users.User, token string) UserResponse {
	format := UserResponse{
		ID:    Helpers.Encrypt(strconv.Itoa(user.ID)),
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return format
}

func UserFormatRegular(user Users.User) UserResponseRegular {
	format := UserResponseRegular{
		ID:    Helpers.Encrypt(strconv.Itoa(user.ID)),
		Name:  user.Name,
		Email: user.Email,
	}

	return format
}
