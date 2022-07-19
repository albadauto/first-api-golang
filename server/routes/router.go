package routes

import (
	"github.com/albadauto/projeto/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
		}

		peoples := main.Group("peoples")
		{
			peoples.GET("/", controllers.ShowPeoples)
			peoples.POST("/", controllers.InsertPeoples)
		}
	}
	return router

}
