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

func AddCategory(db *sql.DB) {

	categories := model.Category{}
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&categories)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
		return
	}

	if categories.Name == "" {
		response := model.ResponseCreate{
			StatusCode: 400,
			Message:    "Invalid input: All fields are required",
			Data:       nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", "  ")

		fmt.Println(string(jsonData))
		return
	}

	repo := repository.NewCatgeoryRepository(db)
	categoryService := service.NewCategoryService(repo)

	category, err := categoryService.InputDataCategory(categories.Name)
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
		Data:       category,
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	fmt.Println(string(jsonData))
}