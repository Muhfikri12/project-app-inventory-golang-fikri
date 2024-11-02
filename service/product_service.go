package service

import (
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

func (ps *ProductService) UpdateDataProduct(product *model.Products) error {

	err := ps.RepoProduct.UpdateProduct(product)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}

func (ps *ProductService) GetDataProducts(pageNumber, pageSize int) ([]model.Products, error) {
	
	products, err := ps.RepoProduct.GetAllDataProducts(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	return products, nil
}