package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

// Configuration of  database
var urlDSN = "root:password@tcp(localhost:3306)/custinfo?parseTime=true"
var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Connection failed")
	}
	Database.AutoMigrate(&Customer{})
}
