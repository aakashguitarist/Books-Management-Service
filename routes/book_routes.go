package routes

import (
	"books-management/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBookByID)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}
}
