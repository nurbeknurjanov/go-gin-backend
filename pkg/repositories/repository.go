package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
)

type Users interface {
	Create(*models.User) error
	Update(*models.User, *models.UserPartial) error
	Delete(*models.User) error
	Find(string) (*models.User, error)
	List(*PaginationRequest, *Sort, *models.UserFilter) ([]*models.User, error)
	Count(*models.UserFilter) (*int, error)
	FindByEmail(string) (*models.User, error)
	ChangeUserPassword(u *models.User, password string) error
}

type Products interface {
	Create(*models.Product) error
	CreateWithImage(*models.Product, *models.File, Files) error
	Update(*models.Product, *models.ProductPartial) error
	Delete(*models.Product) error
	Find(string) (*models.Product, error)
	List(*PaginationRequest, *Sort, *models.ProductFilter) ([]*models.Product, error)
	Count(*models.ProductFilter) (*int, error)
}

type Files interface {
	Create(*models.File) error
	List(*PaginationRequest, *Sort, *models.FileFilter) ([]*models.File, error)
	Count(*models.FileFilter) (*int, error)
	Delete(*models.File) error
	Find(string) (*models.File, error)
	Update(*models.File, *models.FilePartial, *sql.Tx) error
}

// agnostic Repositories
type Repositories struct {
	//Db *sqlx.DB
	Users
	Products
	Files
}

// detailed
func NewSqlRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		//Db:       db,
		Users:    newUsersSqlRepository(db),
		Products: newProductsSqlRepository(db),
		Files:    newFilesSqlRepository(db),
	}
}
