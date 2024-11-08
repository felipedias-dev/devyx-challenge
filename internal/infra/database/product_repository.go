package database

import (
	"database/sql"
	"net/url"
	"strconv"

	"github.com/devyx-tech/felipe-challegend/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Save(product *entity.Product) error {
	stmt, err := r.DB.Prepare("INSERT INTO products (id, name, price) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) GetByID(id string) (*entity.Product, error) {
	stmt, err := r.DB.Prepare("SELECT id, name, price FROM products WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product entity.Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) List(queryParams url.Values) ([]*entity.Product, error) {
	page := queryParams.Get("page")
	if page == "" {
		page = "1"
	}

	limit := queryParams.Get("limit")
	if limit == "" {
		limit = "10"
	}

	pageNumber, _ := strconv.Atoi(page)
	limitNumber, _ := strconv.Atoi(limit)

	offset := (pageNumber - 1) * limitNumber

	stmt, err := r.DB.Prepare("SELECT id, name, price FROM products OFFSET $1 LIMIT $2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(offset, limitNumber)
	if err != nil {
		return nil, err
	}

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (r *ProductRepository) Update(product *entity.Product) error {
	stmt, err := r.DB.Prepare("UPDATE products SET name = $1, price = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Delete(id string) error {
	stmt, err := r.DB.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
