package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Database() *gorm.DB {
	dsn := GetEnv("DB_USERNAME") + ":" + GetEnv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + GetEnv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	
	log.Println("Success Connected to database server")

	return db
}
