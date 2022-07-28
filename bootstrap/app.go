package bootstrap

import (
	"gorm.io/gorm"
	"mvcGolang/config"
)

func Start() *gorm.DB {
	return config.Database()
}
