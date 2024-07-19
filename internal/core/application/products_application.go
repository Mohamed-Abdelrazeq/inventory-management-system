package applications

import "github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/core/repositories"

type ProductsApplication interface {
	GetProducts()
	GetProductById()
	CreateProduct()
	UpdateProduct()
}

type ProductsApplicationInstance struct {
	ProductsRepository *repositories.ProductsRepository
}

func NewProductRepository(productsRepository *repositories.ProductsRepository) ProductsApplication {
	return ProductsApplicationInstance{productsRepository}
}

func (hander ProductsApplicationInstance) GetProducts() {

}

func (hander ProductsApplicationInstance) GetProductById() {

}

func (hander ProductsApplicationInstance) CreateProduct() {

}

func (hander ProductsApplicationInstance) UpdateProduct() {

}
