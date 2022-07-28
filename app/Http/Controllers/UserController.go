package Controllers

import (
	"github.com/gin-gonic/gin"
	"mvcGolang/app/Repository"
)

type UserController interface {
	GetAllUser()
	GetById(c *gin.Context)
}

type userController struct {
	userRepository Repository.UserRepository
}

func NewUserController(userRepository Repository.UserRepository) *userController {
	return &userController{userRepository}
}

func (uc *userController) GetAllUser(c *gin.Context) {
	users, err := uc.userRepository.GetAll()

	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, users)
}

func (uc *userController) GetById(c *gin.Context) {
	user, err := uc.userRepository.GetById(c)

	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, user)
}
