package service

import (
	"fmt"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/repository"
)

type CategoryService struct {
	RepoCategory repository.CategoryRepositoryDB
}

func NewCategoryService(repoCategory repository.CategoryRepositoryDB) *CategoryService {
	return &CategoryService{RepoCategory: repoCategory}
}

func (ps *CategoryService) InputDataCategory(name string) (*model.Category, error) {

	category := &model.Category {
		Name: name,
	}

	err := ps.RepoCategory.CreateCategory(category)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Berhasil tambah data category dengan id", category.ID)

	return category, nil

}