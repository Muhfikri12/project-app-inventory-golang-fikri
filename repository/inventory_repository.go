package repository

import (
	"database/sql"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
)

type InventoryRepositoryDB struct {
	DB *sql.DB
}

func NewInventoryRepository(db *sql.DB) InventoryRepositoryDB {
	return InventoryRepositoryDB{DB: db}
}

func (i *InventoryRepositoryDB) CreateInventory(inventory *model.Inventory) error {
	
	query := `INSERT INTO inventories (product_id, row, part) VALUES ($1, $2, $3) RETURNING ID`

	err := i.DB.QueryRow(query, inventory.ProductId, inventory.Row, inventory.Part).Scan(&inventory.ID)
	if err != nil {
		return err
	}

	return nil
}

func (i *InventoryRepositoryDB) UpdateInventory(inventory *model.Inventory) error {
	
	query := `UPDATE inventories SET product_id=$1, row=$2, part=$3 WHERE product_id=$4`

	_, err := i.DB.Exec(query, inventory.ProductId, inventory.Row, inventory.Part, inventory.ProductId)
	if err != nil {
		return err
	}

	return nil
}


func (i *InventoryRepositoryDB) ChectExistData(ProductId int) (bool, error) {
	
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM inventories WHERE product_id=$1)`
	err := i.DB.QueryRow(query, ProductId).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}