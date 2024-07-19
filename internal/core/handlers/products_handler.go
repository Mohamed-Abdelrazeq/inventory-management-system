package handlers

import (
	applications "github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/core/application"
	"github.com/gin-gonic/gin"
)

type ProductsHandler interface {
	GetProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

type ProductsHandlerInstance struct {
	productsApp *applications.ProductsApplication
}

func NewProductHander(productsApp *applications.ProductsApplication) ProductsHandler {
	return ProductsHandlerInstance{productsApp}
}

func (hander ProductsHandlerInstance) GetProducts(ctx *gin.Context) {
	ctx.JSON(
		201,
		gin.H{
			"message": "working",
		},
	)
}

func (hander ProductsHandlerInstance) GetProductById(ctx *gin.Context) {
	ctx.JSON(
		202,
		gin.H{
			"message": "working",
		},
	)
}

func (hander ProductsHandlerInstance) CreateProduct(ctx *gin.Context) {
	ctx.JSON(
		200,
		gin.H{
			"message": "working",
		},
	)
}

func (hander ProductsHandlerInstance) UpdateProduct(ctx *gin.Context) {
	ctx.JSON(
		200,
		gin.H{
			"message": "working",
		},
	)
}
