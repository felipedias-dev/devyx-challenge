package entity

import "net/url"

type ProductRepositoryInterface interface {
	Save(product *Product) error
	List(queryParams url.Values) ([]*Product, error)
	GetByID(id string) (*Product, error)
	Update(product *Product) error
	Delete(id string) error
}
