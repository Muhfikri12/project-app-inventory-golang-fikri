package repository

import (
	"database/sql"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
)

type ProductRepositoryDB struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryDB {
	return ProductRepositoryDB{DB: db}
}

func (p *ProductRepositoryDB) CreateProduct(product *model.Products) error {
	query := `INSERT INTO products (name, code, stocks, category_id) VALUES ($1, $2, $3, $4) RETURNING ID`

	err := p.DB.QueryRow(query, product.Name, product.Code, product.Stocks, product.CategoryID).Scan(&product.ID)
	if err != nil {
		return err
	}

	return nil
}