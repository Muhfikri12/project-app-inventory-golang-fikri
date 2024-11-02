package repository

import (
	"database/sql"

	"github.com/Muhfikri12/project-app-inventory-golang-fikri/model"
)

type RepositoryUserDB struct {
	DB *sql.DB
}

func NewRepositoryUser(db *sql.DB) RepositoryUserDB  {
	return RepositoryUserDB{DB: db}
}

func (u *RepositoryUserDB) UserLogin(user model.Users) (*model.Users, error) {
	query := `SELECT id, username, password, status
			FROM users
			WHERE username=$1
			AND password=$2`

	var userRespone model.Users

	err := u.DB.QueryRow(query, user.Username, user.Password).Scan(&userRespone.ID, &userRespone.Username, &userRespone.Password, &userRespone.Status)

	if err != nil {
		return nil, err
	}

	return &userRespone, nil
}

func (u *RepositoryUserDB) UpdateStatus(user *model.Users) error {
	query := "UPDATE users SET status =$1 WHERE id =$2"
	_, err := u.DB.Exec(query, user.Status, user.ID)
	return err
}

func (u *RepositoryUserDB) HasActiveUser() (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE status = true)`
	err := u.DB.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *RepositoryUserDB) UserLogout(user model.Users) (*model.Users, error) {
	query := `SELECT id, username, password, status
			FROM users
			WHERE username=$1
			AND password=$2`

	var userRespone model.Users

	err := u.DB.QueryRow(query, user.Username, user.Password).Scan(&userRespone.ID, &userRespone.Username, &userRespone.Password, &userRespone.Status)

	if err != nil {
		return nil, err
	}

	return &userRespone, nil
}

