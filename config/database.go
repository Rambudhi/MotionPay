package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *sql.DB
var Gorm *gorm.DB

func InitializeDB() *gorm.DB {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databaseName, port)
	Database, err = sql.Open("postgres", dbInfo)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	gorm, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the gorm")
	}

	Gorm = gorm

	return gorm
}
