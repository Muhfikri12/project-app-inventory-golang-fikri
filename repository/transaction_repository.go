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

func (t *TransactionRepositoryDB) DeleteTransaction(id int) error{
	
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

func (p *TransactionRepositoryDB) GetAllDataTransaction(page, limit int) ([]model.Transaction, error) {
	offset := (page - 1) * limit
	query := `SELECT id, product_id, qty, is_out FROM transactions LIMIT $1 OFFSET $2`

	rows, err := p.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.ID, &transaction.ProductId, &transaction.Qty, &transaction.IsOut)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}


func (p *TransactionRepositoryDB) CountTotalItems() (int, error) {
	var totalItems int
	query := `SELECT COUNT(*) FROM transactions`
	err := p.DB.QueryRow(query).Scan(&totalItems)
	if err != nil {
		return 0, err
	}
	return totalItems, nil
}