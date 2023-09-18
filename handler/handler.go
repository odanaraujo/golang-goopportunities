package handler

import (
	"github.com/odanaraujo/gopportunities/config"
	"github.com/odanaraujo/gopportunities/database"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = database.GetLogger("handler")
	db = database.GetDB()
}
