package main

import (
	"github.com/odanaraujo/gopportunities/config"
	"github.com/odanaraujo/gopportunities/database"
	"github.com/odanaraujo/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = database.GetLogger("main")

	if err := database.Init(); err != nil {
		logger.Errorf("unable to initialize database: %v", err)
		return
	}

	router.Initialize()
}
