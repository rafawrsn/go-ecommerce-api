package main

import (
	"go-ecommerce-api/config"
	"go-ecommerce-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()

	routes.CategoryRoutes(router)
	routes.ProductRoutes(router)

	router.Run(":8090")
}