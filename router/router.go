package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//initialize Router
	router := gin.Default()

	//initialize Routes
	InitializeRoutes(router)

	//run the server
	router.Run(":8080")
}
