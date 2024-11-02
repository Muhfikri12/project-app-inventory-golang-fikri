package main

import (
	"fmt"
	"log"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/database"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/handler"
	_ "github.com/lib/pq"
)

func main() {

	db, err := database.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var endpoint string
	fmt.Print("masukkan endpoint : ")
	fmt.Scan(&endpoint)

	switch endpoint {
	case "login":
		handler.Login(db)
	case "add/product":
		handler.AddProduct(db)
	case "update/product":
		handler.UpdateProduct(db)
	case "products":
		handler.GetProducts(db)
	case "add/transaction":
		handler.AddTransaction(db)
	case "add/inventory":
		handler.Inventory(db)
	case "delete/transaction":
		handler.DeleteTransaction(db)
	}
}