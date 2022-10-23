package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open("mysql", "root:donnuts54@tcp(127.0.0.1:3306)/empleados?charset=utf8&parseTime=True")

	// database, err := gorm.Open(mysql.Open(sqlinfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Empleado{})

	DB = database
}
