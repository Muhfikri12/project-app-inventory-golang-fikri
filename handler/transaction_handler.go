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
	// Buka file JSON yang berisi data transaksi
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error opening body.json:", err)
		return
	}
	defer file.Close()

	var transactionData model.Transaction

	// Decode JSON dari file ke struct `Transaction`
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&transactionData)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Inisialisasi repository dan service
	repoTransaction := repository.NewTransactionRepository(db)
	productRepo := repository.NewProductRepository(db)
	transactionService := service.NewTransactionService(repoTransaction, productRepo)

	// Panggil service untuk membuat transaksi
	transaction, err := transactionService.InputDataTransaction(transactionData.ProductId, transactionData.Qty, transactionData.IsOut)
	if err != nil {
		response := model.Response{
			StatusCode: 400,
			Message:    "Transaction failed: " + err.Error(),
			Data:       nil,
		}

		jsonData, _ := json.MarshalIndent(response, "", " ")
		fmt.Println(string(jsonData))
		return
	}

	// Respon sukses dengan data transaksi
	response := model.Response{
		StatusCode: 200,
		Message:    "Transaction successful",
		Data:       transaction,
	}

	jsonData, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(jsonData))
}