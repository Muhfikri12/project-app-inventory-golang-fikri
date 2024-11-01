package service

import (
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
	"github.com/Muhfikri12/project-app-inventory-golang-fikri/repository"
)

type UserService struct {
	RepoUser repository.RepositoryUserDB
}

func NewUserService(repo repository.RepositoryUserDB) *UserService {
	return &UserService{RepoUser: repo}
}

func (us *UserService) LoginService(user model.Users) (*model.Users, error) {
	
	users, err := us.RepoUser.UserLogin(user)

	if err != nil {
		return nil, err
	}

	users.Status = true

	err = us.RepoUser.UpdateStatus(users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
