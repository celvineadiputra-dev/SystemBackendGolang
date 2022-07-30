package bootstrap

import (
	"gorm.io/gorm"
	"mvcGolang/config"
	"mvcGolang/database"
)

func Start() *gorm.DB {
	return config.Database()
}

func MigrateDatabase() {
	database.Migrate()
}
