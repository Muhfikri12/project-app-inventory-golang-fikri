package service

import (
	"errors"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/repository"
)

type InventoryService struct {
	RepoInventory repository.InventoryRepositoryDB
}

func NewInventoryService(repo repository.InventoryRepositoryDB) *InventoryService {
	return &InventoryService{RepoInventory: repo}
}


func (is *InventoryService) InputDataInventory(productId, row, part int) (*model.Inventory, error) {
	
	inventory := &model.Inventory{
		ProductId: productId,
		Row: row,
		Part: part,
	}

	exists, err := is.RepoInventory.ChectExistData(productId)
	if err != nil {
		return nil, err
	}

	if exists {
		err = is.RepoInventory.UpdateInventory(inventory)
		if err != nil {
			return nil, errors.New("failed to update inventory: " + err.Error())
		}
	} else {
		err = is.RepoInventory.CreateInventory(inventory)
		if err != nil {
			return nil, errors.New("failed to create inventory: " + err.Error())
		}
	}

	return inventory, nil
}