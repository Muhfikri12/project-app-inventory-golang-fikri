package service

import (
	"errors"
	"fmt"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/repository"
)

type TransactionService struct {
	RepoTransaction repository.TransactionRepositoryDB
	RepoProduct repository.ProductRepositoryDB
}

func NewTransactionService(repoTransaction repository.TransactionRepositoryDB, repoProduct repository.ProductRepositoryDB) *TransactionService {
	return &TransactionService{
		RepoTransaction: repoTransaction,
		RepoProduct: repoProduct,
	}
	
}

func NewTransactionServiceDelete(DeleteTransaction repository.TransactionRepositoryDB) *TransactionService {
	return &TransactionService{RepoTransaction: DeleteTransaction}
}

func (ts *TransactionService) InputDataTransaction(productID, qty int, isOut bool) (*model.Transaction, error) {
	product, err := ts.RepoProduct.GetProductByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if isOut {
		if product.Stocks < qty {
			return nil, errors.New("insufficient stock")
		}
		product.Stocks -= qty
	} else {
		product.Stocks += qty
	}

	transaction := &model.Transaction{
		ProductId: productID,
		Qty:       qty,
		IsOut:     isOut,
	}

	err = ts.RepoTransaction.CreateTransaction(transaction, &product)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return transaction, nil
}

func (ts *TransactionService) DeletingTransaction(id int) ( error) {
	
	exists, err := ts.RepoTransaction.ChectExistsData(id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("transaction not found")
	}

	err = ts.RepoTransaction.DeleteTransaction(id)
	if err != nil {
		return errors.New("failed to delete transaction: " + err.Error())
	}

	return nil
}

func (ts *TransactionService) GetDataTransactions(page, limit int) (int, int, []model.Transaction, error) {
	totalItems, err := ts.RepoTransaction.CountTotalItems()
	if err != nil {
		return 0, 0, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit

	transactions, err := ts.RepoTransaction.GetAllDataTransaction(page, limit)
	if err != nil {
		return 0, 0, nil, err
	}

	return totalItems, totalPages, transactions, nil
}