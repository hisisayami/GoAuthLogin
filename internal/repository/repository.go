package repository

import (
	"goauthlogin/internal/models"
	"database/sql"
	
)

type DatabaseRepo interface {
	Connection() *sql.DB
	CreateUserTable() error
	GetUserName(username string) (*models.User, error)
	CreateUser(user *models.User) error
}
