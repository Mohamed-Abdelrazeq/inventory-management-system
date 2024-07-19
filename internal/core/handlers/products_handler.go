package handlers

import (
	"net/http"

	applications "github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/core/application"
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/database"
	"github.com/gin-gonic/gin"
)

type ProductsHandler interface {
	GetProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

type ProductsHandlerInstance struct {
	productsApp applications.ProductsApplication
}

func NewProductHander(productsApp *applications.ProductsApplication) ProductsHandler {
	return ProductsHandlerInstance{*productsApp}
}

func (handler ProductsHandlerInstance) GetProducts(ctx *gin.Context) {
	params := database.GetProductsByLocationIdParams{
		LocationID: 1,
		Limit:      5,
		Offset:     0,
	}
	data, err := handler.productsApp.GetProducts(params)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err,
			},
		)
		return
	}
	ctx.JSON(
		http.StatusAccepted,
		gin.H{
			"data": data,
		},
	)
}

func (handler ProductsHandlerInstance) GetProductById(ctx *gin.Context) {
	ctx.JSON(
		202,
		gin.H{
			"message": "working",
		},
	)
}

func (handler ProductsHandlerInstance) CreateProduct(ctx *gin.Context) {
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
