package applications

import (
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/core/repositories"
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/database"
)

type ProductsApplication interface {
	GetProducts(database.GetProductsByLocationIdParams) ([]database.GetProductsByLocationIdRow, error)
	GetProductById()
	CreateProduct()
	UpdateProduct()
}

type ProductsApplicationInstance struct {
	ProductsRepository repositories.ProductsRepository
}

func NewProductRepository(productsRepository *repositories.ProductsRepository) ProductsApplication {
	return ProductsApplicationInstance{*productsRepository}
}

func (app ProductsApplicationInstance) GetProducts(params database.GetProductsByLocationIdParams) ([]database.GetProductsByLocationIdRow, error) {
	return app.ProductsRepository.GetProducts(params)
}

func (app ProductsApplicationInstance) GetProductById() {

}

func (app ProductsApplicationInstance) CreateProduct() {

}

func (app ProductsApplicationInstance) UpdateProduct() {

}
