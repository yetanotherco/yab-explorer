package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

const (
	DATABASE_PORT = "5432"
)

func ConnectToDB() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	if port == "" {
		port = DATABASE_PORT
	}

	databaseConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, name, port)

	db, err := gorm.Open(postgres.Open(databaseConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	dbInstance = db
	return dbInstance
}
