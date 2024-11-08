package entity

import "github.com/google/uuid"

type Product struct {
	ID    string  `sql:"id" json:"id"`
	Name  string  `sql:"name" json:"name"`
	Price float64 `sql:"price" json:"price"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
