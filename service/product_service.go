package service

import (
	"errors"
	"fmt"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/repository"
)


type ProductService struct {
	RepoProduct repository.ProductRepositoryDB
}

func NewProductService(repoProduct repository.ProductRepositoryDB) *ProductService {
	return &ProductService{RepoProduct: repoProduct}
}

func (ps *ProductService) InputDataProduct(name, code string, stocks, categoryId int) error {
	if name == "" {
		return errors.New("name is required")
	}

	if code == "" {
		return errors.New("code is required")
	}

	if stocks == 0 {
		return errors.New("value should be more than 0")
	}

	product := model.Products {
		Name: name,
		Code: code,
		Stocks: stocks,
		CategoryID: categoryId,
	}

	err := ps.RepoProduct.CreateProduct(&product)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Berhasil tambah data Produk dengan id", product.ID)

	return nil

}