package repositories

import "github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/loaders"

type ProductsRepository interface {
	GetProducts()
	GetProductById()
	CreateProduct()
	UpdateProduct()
}

type ProductsRepositoryInstance struct {
	DB *loaders.DatabaseInstance
}

func NewProductRepository(db *loaders.DatabaseInstance) ProductsRepository {
	return ProductsRepositoryInstance{db}
}

func (hander ProductsRepositoryInstance) GetProducts() {

}

func (hander ProductsRepositoryInstance) GetProductById() {

}

func (hander ProductsRepositoryInstance) CreateProduct() {

}

func (hander ProductsRepositoryInstance) UpdateProduct() {

}
