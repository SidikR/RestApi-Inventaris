package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db, err := gorm.Open("mysql", databaseURL())
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

func databaseURL() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
}
