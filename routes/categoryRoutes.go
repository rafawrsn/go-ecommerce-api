package routes

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/handlers"
)

func CategoryRoutes(router *gin.Engine) {
	router.POST("/categories", handlers.CreateCategory)
	router.GET("/categories", handlers.GetCategories)
	router.GET("/categories/:id", handlers.GetCategoryByID)
	router.PUT("/categories/:id", handlers.UpdateCategory)
	router.DELETE("/categories/:id", handlers.DeleteCategory)
}