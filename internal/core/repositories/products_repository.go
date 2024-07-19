package repositories

import (
	"context"
	"fmt"

	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/database"
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/loaders"
)

type ProductsRepository interface {
	GetProducts(database.GetProductsByLocationIdParams) ([]database.GetProductsByLocationIdRow, error)
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

func (repo ProductsRepositoryInstance) GetProducts(params database.GetProductsByLocationIdParams) ([]database.GetProductsByLocationIdRow, error) {
	data, err := repo.DB.DB.GetProductsByLocationId(context.TODO(), params)
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func (repo ProductsRepositoryInstance) GetProductById() {

}

func (repo ProductsRepositoryInstance) CreateProduct() {

}

func (repo ProductsRepositoryInstance) UpdateProduct() {

}
