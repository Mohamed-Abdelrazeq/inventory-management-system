package repositories

import (
	"context"
	"log"

	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/database"
	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/loaders"
)

type ProductsRepository interface {
	GetProducts(database.GetProductsByLocationIdParams) ([]database.GetProductsByLocationIdRow, error)
	GetProductById(id int32) (database.Product, error)
	CreateProduct(database.CreateProductParams) (database.Product, error)
	UpdateProduct(database.UpdateProductParams) (database.Product, error)
}

type ProductsRepositoryInstance struct {
	DB *loaders.DatabaseInstance
}

func NewProductRepository(db *loaders.DatabaseInstance) ProductsRepository {
	return ProductsRepositoryInstance{db}
}

func (repo ProductsRepositoryInstance) GetProducts(params database.GetProductsByLocationIdParams) ([]database.GetProductsByLocationIdRow, error) {
	products, err := repo.DB.DB.GetProductsByLocationId(context.TODO(), params)
	if err != nil {
		log.Println(err)
		return products, err
	}
	return products, nil
}

func (repo ProductsRepositoryInstance) GetProductById(id int32) (database.Product, error) {
	product, err := repo.DB.DB.GetProductById(context.TODO(), id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	return product, nil
}

func (repo ProductsRepositoryInstance) CreateProduct(params database.CreateProductParams) (database.Product, error) {
	product, err := repo.DB.DB.CreateProduct(context.TODO(), params)
	if err != nil {
		log.Println(err)
		return product, err
	}
	return product, nil

}

func (repo ProductsRepositoryInstance) UpdateProduct(params database.UpdateProductParams) (database.Product, error) {
	product, err := repo.DB.DB.UpdateProduct(context.TODO(), params)
	if err != nil {
		log.Println(err)
		return product, err
	}
	return product, nil
}
