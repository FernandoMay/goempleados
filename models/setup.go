package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:donnuts64@tcp(127.0.0.1:3306)/empleados?charset=utf8&parseTime=True&loc=Local"

func ConnectDatabase() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	// database, err := gorm.Open("mysql", "sammy:password@tcp(127.0.0.1:3306)/empleados.db?charset=utf8&parseTime=True")

	// database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&Empleado{})
}
