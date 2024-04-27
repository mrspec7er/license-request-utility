package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	CreateDB()
}

func CreateDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable TimeZone=Asia/Singapore", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	cmd := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", os.Getenv("DB_NAME"))
	db.Exec(cmd)

	fmt.Println("Database created!")

	return db
}
