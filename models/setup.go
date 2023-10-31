package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/gorm"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Murid{})

	DB = database
}