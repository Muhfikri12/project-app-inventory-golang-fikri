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


func Inventory(db *sql.DB)  {
	
	inventories := model.Inventory{}

	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error Opening body.json", err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&inventories)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
		return
	}

	if inventories.ProductId <= 0 || inventories.Row <= 0 || inventories.Part <= 0 {
		response := model.ResponseCreate {
			StatusCode: 400,
			Message: "Invalid input: All field is required",
			Data: nil,
		}
		jsonData, _ := json.MarshalIndent(response, "", " ")
		fmt.Println(string(jsonData))
		return
	}

	repo := repository.NewInventoryRepository(db)
	service := service.NewInventoryService(repo)

	product, err := service.InputDataInventory(inventories.ProductId, inventories.Row, inventories.Part)
	if err != nil {
		response := model.ResponseCreate {
			StatusCode: 400,
			Message: "Error Adding or Update Inventory",
			Data: nil,
		}

		jsonData, _ := json.MarshalIndent(response, "", " ")
		fmt.Println(string(jsonData))
		return
	}

	response := model.ResponseCreate {
		StatusCode: 201,
		Message: "Successfully Added inventory",
		Data: product,
	}

	jsonData, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(jsonData))

}