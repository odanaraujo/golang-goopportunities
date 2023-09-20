package router

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/gopportunities/handler"
)

func InitializeRoutes(route *gin.Engine) {
	handler.InitializeHandler()
	v1 := route.Group("/api/v1")
	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
	}
}
