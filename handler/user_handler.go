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

var CurrentUserID int

func Login(db *sql.DB)  {
	
	inputUser := model.Users{}

	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("error: ", err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&inputUser)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON: ", err)
	}

	repo := repository.NewRepositoryUser(db)
	loginService := service.NewUserService(repo)

	user, err := loginService.LoginService(inputUser)

	if err != nil {
		response := model.Response{
			StatusCode: 404,
			Message:    "Account not found",
			Data:       nil,
		}
		jsonData, err := json.MarshalIndent(response, " ", " ")

		if err != nil {
			fmt.Println("err :", err)
		}

		fmt.Println(string(jsonData))
	} else {
		CurrentUserID = user.ID
		
		response := model.Response{
			StatusCode: 200,
			Message:    "login success",
			Data:       user,
		}
		jsonData, err := json.MarshalIndent(response, " ", "")

		if err != nil {
			fmt.Println("err :", err)
		}

		fmt.Println(string(jsonData))
	}
}