// handler/product_handler.go

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

func AddProduct(db *sql.DB) {

	products := model.Products{}
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
		return
	}

	if products.Name == "" || products.Code == "" || products.Stocks == 0 || products.CategoryID == 0 {
		response := model.Response{
			StatusCode: 400,
			Message:    "Invalid input: All fields are required",
			Data:       nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", "  ")
		
		fmt.Println(string(jsonData))
		return
	}

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	product := productService.InputDataProduct(products.Name, products.Code, products.Stocks, products.CategoryID)
	if err != nil {
		response := model.Response{
			StatusCode: 400,
			Message:    "Error adding product",
			Data:       nil,
		}

		jsonData, err := json.MarshalIndent(response, "", " ")
		if err != nil {
			fmt.Println("err :", err)
			return
		}

		fmt.Println(string(jsonData))
		return
	}

	response := model.Response{
		StatusCode: 200,
		Message:    "Product added successfully",
		Data:       product,
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	fmt.Println(string(jsonData))
}
