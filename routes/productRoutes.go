package routes

import (
	"go-ecommerce-api/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.POST("/products", handlers.CreateProduct)
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProductByID)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)
}