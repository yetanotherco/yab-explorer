package configs

import (
	"fmt"
	"yab-explorer/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func NewDBConnection(host, user, password, dbName, port string) *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&models.Order{})

	if err != nil {
		panic("Failed to migrate database")
	}

	dbInstance = db
	return dbInstance
}
