package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/odanaraujo/gopportunities/docs"
	"github.com/odanaraujo/gopportunities/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(route *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := route.Group(basePath)
	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
