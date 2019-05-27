package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateConnection() (*gorm.DB, error) {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, DBName)
	return gorm.Open("mysql", connectionString)

	// fmt.Sprintf(
	// 	"postgres://%s:%s@%s/%s?sslmode=disable",
	// 	user, password, host, DBName,
	// ),

}
