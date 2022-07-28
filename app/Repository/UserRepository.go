package Repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Models/Users"
)

type UserRepository interface {
	GetAll() ([]Users.User, error)
	GetById(c *gin.Context) (Users.User, error)
	RegisterUser(user Users.User) (Users.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]Users.User, error) {
	var users []Users.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) GetById(c *gin.Context) (Users.User, error) {
	var user Users.User

	err := r.db.Where("id", c.Param("id")).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) RegisterUser(user Users.User) (Users.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
