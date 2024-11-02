package repository

import (
	"database/sql"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
)

type TransactionRepositoryDB struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{DB: db}
}

func (t *TransactionRepositoryDB) CreateTransaction(transaction *model.Transaction, product *model.Products) error{
	tx, err := t.DB.Begin()

	defer func ()  {
		if err != nil {
			tx.Rollback()
		}
	}()

	transactionQuery := `INSERT INTO transactions (product_id, qty, is_out) VALUES($1, $2, $3) RETURNING ID`

	err = tx.QueryRow(transactionQuery, transaction.ProductId, transaction.Qty, transaction.IsOut).Scan(&transaction.ID)
	if err != nil {
		return err
	}

	productQuery := `UPDATE products SET stocks=$1 WHERE id=$2`

	_, err = tx.Exec(productQuery, product.Stocks, product.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t *TransactionRepositoryDB) DeleteTransaction (id int) error{
	
	query := `DELETE FROM transactions WHERE id=$1`

	_, err := t.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (i *TransactionRepositoryDB) ChectExistsData(id int) (bool, error) {
	
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM transactions WHERE id=$1)`
	err := i.DB.QueryRow(query, id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}