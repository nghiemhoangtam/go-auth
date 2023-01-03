package database

import (
	"goauth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:12345@tcp(localhost:3306)/goauth?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})

}
