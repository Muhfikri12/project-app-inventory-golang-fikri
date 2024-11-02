package repository

import (
	"database/sql"
	"errors"
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


func (p *ProductRepositoryDB) GetAllDataProducts(page, limit int) ([]model.Pagination,[]model.Products, error) {
	offset := (page - 1) * limit
	query := `SELECT id, name, code, stocks, category_id FROM products LIMIT $1 OFFSET $2`

	rows, err := p.DB.Query(query, limit, offset)

	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	var pagination []model.Pagination

	var products []model.Products

	for rows.Next() {
		var product model.Products
		err := rows.Scan(&product.ID, &product.Name,&product.Code, &product.Stocks, &product.CategoryID)
		if err != nil {
			return nil, nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
        return nil, nil, err
    }



	return pagination, products, nil
}

func (p *ProductRepositoryDB) CountTotalItems() (int, error) {
	var totalItems int
	query := `SELECT COUNT(*) FROM products`
	err := p.DB.QueryRow(query).Scan(&totalItems)
	if err != nil {
		return 0, err
	}
	return totalItems, nil
}

func (p *ProductRepositoryDB) GetProductByID(productID int) (model.Products, error) {
	query := `SELECT id, name, code, stocks, category_id FROM products WHERE id = $1`

	var product model.Products
	err := p.DB.QueryRow(query, productID).Scan(&product.ID, &product.Name, &product.Code, &product.Stocks, &product.CategoryID)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, errors.New("product not found")
		}
		return product, err
	}

	return product, nil
}

