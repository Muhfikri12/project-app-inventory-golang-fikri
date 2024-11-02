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

func (ps *ProductService) InputDataProduct(name, code string, stocks, categoryId int) (*model.Products, error) {

	product := &model.Products {
		Name: name,
		Code: code,
		Stocks: stocks,
		CategoryID: categoryId,
	}

	err := ps.RepoProduct.CreateProduct(product)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Berhasil tambah data Produk dengan id", product.ID)

	return product, nil

}

func (ps *ProductService) UpdateDataProduct(product *model.Products) error {

	err := ps.RepoProduct.UpdateProduct(product)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}

func (ps *ProductService) GetDataProducts(page, limit int) (int, int, []model.Products, error) {
	totalItems, err := ps.RepoProduct.CountTotalItems()
	if err != nil {
		return 0, 0, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit

	_, products, err := ps.RepoProduct.GetAllDataProducts(page, limit)
	if err != nil {
		return 0, 0, nil, err
	}

	return totalItems, totalPages, products, nil
}


func (ts *ProductService) DeletingProduct(id int) ( error) {
	
	exists, err := ts.RepoProduct.ChectExistsData(id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("product not found")
	}

	err = ts.RepoProduct.DeleteProduct(id)
	if err != nil {
		return errors.New("failed to delete product: " + err.Error())
	}

	return nil
}
