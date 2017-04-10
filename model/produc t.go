package model

import "time"

// swagger:model
type Product struct {
	// Ürün  kimliği
	ID int64 `json:"id"`

	// Ürün ismi
	// min length: 3
	// required: true
	Name string `json:"name"`

	// Ürün fiyatı
	// minimum: 1
	Price float64 `json:"price"`

	// Ürünün üretim tarihi
	// swagger:strfmt date
	ProductionDate time.Time `json:"productionDate"`
}

// swagger:parameters create-product
type ProductInput struct {
	// Eklenecek ürüne ait bilgiler
	// in:body
	Product Product `json:"product"`
}

// swagger:parameters get-product-by-id
type ProductByIdInput struct {
	// Ürün referansı
	// in:path
	Id int64 `json:"id"`
}
