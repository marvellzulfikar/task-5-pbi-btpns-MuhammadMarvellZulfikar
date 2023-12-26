package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/project-btpn"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Photos{})

	DB = database
}
