package database

import (
	"fmt"
	"github.com/odanaraujo/gopportunities/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func Init() error {
	var err error
	db, err = InitializeSqLite()

	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(prefix string) *config.Logger {
	logger = config.NewLogger(prefix)
	return logger
}
