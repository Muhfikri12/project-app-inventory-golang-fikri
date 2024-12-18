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

	if products.Name == "" || products.Code == "" || products.Stocks <= 0 || products.CategoryID == 0 {
		response := model.ResponseCreate{
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

	product, err := productService.InputDataProduct(products.Name, products.Code, products.Stocks, products.CategoryID)
	if err != nil {
		response := model.ResponseCreate{
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

	response := model.ResponseCreate{
		StatusCode: 201,
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
		response := model.ResponseCreate{
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

	err = productService.UpdateDataProduct(&product, product.ID)
	if err != nil {
		response := model.ResponseCreate{
			StatusCode: 400,
			Message:    err.Error(),
			Data:       nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", "  ")
		fmt.Println(string(jsonData))
		return
	}

	response := model.ResponseCreate{
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

	page := pagination.Page
	if page == 0 {
		page = 1
	}

	limit := pagination.Limit
	if limit <= 10 {
		limit = 10
	}

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	totalItems, totalPages, products, err := productService.GetDataProducts(page, limit)
	if err != nil {
		response := model.Response{
			StatusCode: 500,
			Message:    "Error fetching Data Products",
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
		Data:       products,
	}
	jsonData, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(jsonData))
}

func DeleteProduct(db *sql.DB) {
    var product model.Products

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
    err = decoder.Decode(&product)
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

    if product.ID <= 0 {
        response := model.ResponseCreate{
            StatusCode: 400,
            Message:    "Invalid input: id is required",
            Data:       nil,
        }
        jsonData, _ := json.MarshalIndent(response, "", " ")
        fmt.Println(string(jsonData))
        return
    }

    repo := repository.NewProductRepository(db)
    service := service.NewProductService(repo)

    err = service.DeletingProduct(product.ID) 
    if err != nil {
        response := model.ResponseCreate{
            StatusCode: 400,
            Message:    "Error deleting product: " + err.Error(),
            Data:       nil,
        }
        jsonData, _ := json.MarshalIndent(response, "", " ")
        fmt.Println(string(jsonData))
        return
    }

    response := model.ResponseCreate{
        StatusCode: 200,
        Message:    "Successfully deleted product",
        Data:       nil,
    }

    jsonData, _ := json.MarshalIndent(response, "", " ")
    fmt.Println(string(jsonData))
}

func FilterProducts(db *sql.DB) {
    var product struct {
        ID int `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Stocks int `json:"stocks"`
        CategoryID *int  `json:"category_id"`
    }

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
    err = decoder.Decode(&product)
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

    repo := repository.NewProductRepository(db)
    productService := service.NewProductService(repo)

    products, err := productService.FilterProducts(product.Name, product.Code, product.CategoryID)
    if err != nil {
        response := model.ResponseCreate{
            StatusCode: 400,
            Message:    "Error filtering products: " + err.Error(),
            Data:       nil,
        }
        jsonData, _ := json.MarshalIndent(response, "", " ")
        fmt.Println(string(jsonData))
        return
    }

    response := model.ResponseCreate{
        StatusCode: 200,
        Message:    "Successfully retrieved products",
        Data:       products,
    }

    jsonData, _ := json.MarshalIndent(response, "", " ")
    fmt.Println(string(jsonData))
}

func GetProductsless10(db *sql.DB) {
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

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	totalItems, totalPages, products, err := productService.GetDataProductsLess10(page, limit)
	if err != nil {
		response := model.Response{
			StatusCode: 500,
			Message:    "Error fetching Data Products",
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
		Data:       products,
	}
	jsonData, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(jsonData))
}