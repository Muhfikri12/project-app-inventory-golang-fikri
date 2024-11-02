package repository

import (
	"database/sql"
	"time"

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

func (p *ProductRepositoryDB) UpdateProduct(product *model.Products) error { 
	query := `UPDATE products SET name=$1, code=$2, stocks=$3, category_id=$4, updated_at=$5 WHERE id=$6`

	currentTime := time.Now()
	_, err := p.DB.Exec(query, product.Name, product.Code, product.Stocks, product.CategoryID, currentTime, product.ID)
	if err != nil {
		return err
	}

	product.Updated_at = currentTime

	return nil
}