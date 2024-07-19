package main

import (
	applications "github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/core/application"
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/core/handlers"
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/core/repositories"
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/loaders"
	"github.com/gin-gonic/gin"
)

func main() {
	loaders.LoadEnv()
	db := loaders.LoadDB()

	productsRepository := repositories.NewProductRepository(db)
	productsApplication := applications.NewProductRepository(&productsRepository)
	productsHandler := handlers.NewProductHander(&productsApplication)

	// Create Router
	r := gin.Default()

	// Test Routes
	r.GET("/ping", handlers.HealthTest)

	productRouter := r.Group("/products/")
	productRouter.GET(":id", productsHandler.GetProductById)
	productRouter.GET("location/:id", productsHandler.GetProducts)
	productRouter.GET("supplier/:id", productsHandler.GetProducts)
	productRouter.POST("", productsHandler.CreateProduct)
	productRouter.PUT("", productsHandler.UpdateProduct)
	// Run Server
	r.Run()
}
