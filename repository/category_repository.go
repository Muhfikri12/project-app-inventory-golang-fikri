package repository

import (
	"database/sql"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
)

type CategoryRepositoryDB struct {
	DB *sql.DB
}

func NewCatgeoryRepository(db *sql.DB) CategoryRepositoryDB {
	return CategoryRepositoryDB{DB: db}
}

func (c *CategoryRepositoryDB) CreateCategory(category *model.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING ID`

	err := c.DB.QueryRow(query, category.Name).Scan(&category.ID)
	if err != nil {
		return err
	}

	return nil
}