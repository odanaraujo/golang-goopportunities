package database

import (
	"github.com/odanaraujo/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

const (
	dbPath = "./db/main.db"
)

func InitializeSqLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// Check if the database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating a new one!")

		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Error("Failed to create database directory!")
			return nil, err
		}

		file, err := os.Create(dbPath)
		defer file.Close()

		if err != nil {
			logger.Error("Failed to create database file!")
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Error("Failed to connect to database!")
		return nil, err
	}

	if err := db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Error("Failed to migrate database!")
		return nil, err
	}

	logger.Info("Connected to database!")
	return db, nil
}
