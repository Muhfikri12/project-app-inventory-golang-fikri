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

func UpdateProduct(db *sql.DB) {
	product := model.Products{}

	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&product)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if product.ID == 0 {
		response := model.Response{
			StatusCode: 400,
			Message:    "Product ID is required",
			Data:       nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", "  ")
		fmt.Println(string(jsonData))
		return
	}

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	err = productService.UpdateDataProduct(&product)
	if err != nil {
		response := model.Response{
			StatusCode: 400,
			Message:    err.Error(),
			Data:       nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", "  ")
		fmt.Println(string(jsonData))
		return
	}

	response := model.Response{
		StatusCode: 200,
		Message:    "Product updated successfully",
		Data:       product,
	}
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonData))
}

func GetProducts(db *sql.DB) {
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

	pageNumber := pagination.PageNumber
	if pageNumber == 0 {
		pageNumber = 1
	}

	pageSize := pagination.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	products, err := productService.GetDataProducts(pageNumber, pageSize)
	if err != nil {
		response := model.Response{
			StatusCode: 500,
			Message:    "Error fetching Data Products",
			Data:       nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", " ")
		fmt.Println(string(jsonData))
		return
	}

	response := model.Response{
		StatusCode: 200,
		Message:    "Products fetched successfully",
		Data:       products,
	}
	jsonData, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(jsonData))
}