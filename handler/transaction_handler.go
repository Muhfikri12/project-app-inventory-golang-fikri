package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/repository"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/service"
)

func AddTransaction(db *sql.DB) {

	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error opening body.json:", err)
		return
	}
	defer file.Close()

	var transactionData model.Transaction

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&transactionData)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	repoTransaction := repository.NewTransactionRepository(db)
	productRepo := repository.NewProductRepository(db)
	transactionService := service.NewTransactionService(repoTransaction, productRepo)

	transaction, err := transactionService.InputDataTransaction(transactionData.ProductId, transactionData.Qty, transactionData.IsOut)
	if err != nil {
		response := model.ResponseCreate{
			StatusCode: 400,
			Message:    "Transaction failed: " + err.Error(),
			Data:       nil,
		}

		jsonData, _ := json.MarshalIndent(response, "", " ")
		fmt.Println(string(jsonData))
		return
	}

	response := model.ResponseCreate{
		StatusCode: 200,
		Message:    "Transaction successful",
		Data:       transaction,
	}

	jsonData, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(jsonData))
}

func DeleteTransaction(db *sql.DB) {
    var transaction model.Transaction

    file, err := os.Open("body.json")
    if err != nil {
        response := model.ResponseCreate{
            StatusCode: 500,
            Message:    "Error opening body.json: " + err.Error(),
            Data:       nil,
        }
        jsonData, _ := json.MarshalIndent(response, "", " ")
        fmt.Println(string(jsonData))
        return
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&transaction)
    if err != nil && err != io.EOF {
        response := model.ResponseCreate{
            StatusCode: 400,
            Message:    "Error decoding JSON: " + err.Error(),
            Data:       nil,
        }
        jsonData, _ := json.MarshalIndent(response, "", " ")
        fmt.Println(string(jsonData))
        return
    }

    if transaction.ID <= 0 {
        response := model.ResponseCreate{
            StatusCode: 400,
            Message:    "Invalid input: id is required",
            Data:       nil,
        }
        jsonData, _ := json.MarshalIndent(response, "", " ")
        fmt.Println(string(jsonData))
        return
    }

    repo := repository.NewTransactionRepository(db)
    service := service.NewTransactionServiceDelete(repo)

    err = service.DeletingTransaction(transaction.ID) 
    if err != nil {
        response := model.ResponseCreate{
            StatusCode: 400,
            Message:    "Error deleting transaction: " + err.Error(),
            Data:       nil,
        }
        jsonData, _ := json.MarshalIndent(response, "", " ")
        fmt.Println(string(jsonData))
        return
    }

    response := model.ResponseCreate{
        StatusCode: 200,
        Message:    "Successfully deleted transaction",
        Data:       nil,
    }

    jsonData, _ := json.MarshalIndent(response, "", " ")
    fmt.Println(string(jsonData))
}

func GetTransactions(db *sql.DB) {
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error opening body.json:", err)
		return
	}
	defer file.Close()

	var pagination model.Pagination
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pagination)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	page := pagination.Page
	if page == 0 {
		page = 1
	}

	limit := pagination.Limit
	if limit <= 10 {
		limit = 10
	}

	repo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionServiceDelete(repo)

	totalItems, totalPages, transactions, err := transactionService.GetDataTransactions(page, limit)
	if err != nil {
		response := model.Response{
			StatusCode: 500,
			Message:    "Error fetching Data Transactions: " + err.Error(),
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
			Data:       nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", " ")
		fmt.Println(string(jsonData))
		return
	}

	response := model.Response{
		StatusCode: 200,
		Message:    "Data retrieved successfully",
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       transactions,
	}
	jsonData, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(jsonData))
}