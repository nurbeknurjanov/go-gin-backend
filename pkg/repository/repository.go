package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
)

type IUsersRepository interface {
	CreateUser(*models.User) error
	UpdateUser(*models.User, *models.UserPartial) error
	DeleteUser(*models.User) error
	FindUser(string) (*models.User, error)
	ListUsers(*PaginationRequest, *Sort, *models.UserFilter) ([]*models.User, error)
	CountUsers(*models.UserFilter) (*int, error)
	FindByEmail(string) (*models.User, error)

	ChangeUserPassword(u *models.User, password string) error
}

type IProductsRepository interface {
	CreateProduct(*models.Product) error
	CreateProductWithImage(*models.Product, *sql.Tx) error
	UpdateProduct(*models.Product, *models.ProductPartial) error
	DeleteProduct(*models.Product) error
	FindProduct(string) (*models.Product, error)
	ListProducts(*PaginationRequest, *Sort, *models.ProductFilter) ([]*models.Product, error)
	CountProducts(*models.ProductFilter) (*int, error)
}

type IFilesRepository interface {
	CreateFile(*models.File) error
	ListFiles(*PaginationRequest, *Sort, *models.FileFilter) ([]*models.File, error)
	CountFiles(*models.FileFilter) (*int, error)
	DeleteFile(*models.File) error
	FindFile(string) (*models.File, error)
	UpdateFile(*models.File, *models.FilePartial, *sql.Tx) error
}

type Repositories struct {
	Db *sqlx.DB
	IUsersRepository
	IProductsRepository
	IFilesRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Db:                  db,
		IUsersRepository:    newUsersRepository(db),
		IProductsRepository: newProductsRepository(db),
		IFilesRepository:    newFilesRepository(db),
	}
}
