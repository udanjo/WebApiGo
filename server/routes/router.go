package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ueverson/WebApiGo/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		movies := main.Group("asset")
		{
			movies.GET("/:id", controllers.ShowMovie)
			movies.GET("/", controllers.ShowMovies)
			movies.POST("/", controllers.CreateMovie)
			movies.PUT("/", controllers.UpdateMovie)
			movies.DELETE("/", controllers.DeleteMovie)
		}
	}

	return router
}
