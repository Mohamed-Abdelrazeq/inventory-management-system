package models

import (
	_ "github.com/go-playground/validator/v10"
)

type Product struct {
	ProductID   int     `json:"product_id" validate:"required"`
	LocationID  int     `json:"location_id" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	UnitPrice   float64 `json:"unit_price" validate:"required"`
	SupplierID  int     `json:"supplier_id" validate:"required"`
}
