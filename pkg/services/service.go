package services

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
)

const (
	adminEmail = "nurbek.nurjanov@mail.ru"
)

type Auth interface {
	Login(string, string) (*models.Tokens, error)
	GetAccessToken(*models.User) (string, error)
}

type Profile interface {
	UpdateProfile(*models.User, *models.UserPartial) error
}

type Users interface {
	Create(*models.User) error
	Update(*models.User, *models.UserPartial) error
	Delete(*models.User) error
	Find(id string) (*models.User, error)
	List(*repositories.PaginationRequest, *repositories.Sort, *models.UserFilter) ([]*models.User, error)
	Count(*models.UserFilter) (*int, error)
	ChangeUserPassword(*models.User, string) error
}

type Products interface {
	Create(*models.Product) error
	Update(*models.Product, *models.ProductPartial) error
	Delete(*models.Product) error
	Find(id string) (*models.Product, error)
	List(*repositories.PaginationRequest, *repositories.Sort, *models.ProductFilter) ([]*models.Product, error)
	Count(*models.ProductFilter) (*int, error)
}

type Files interface {
	Create(*models.File) error
	List(*repositories.PaginationRequest, *repositories.Sort, *models.FileFilter) ([]*models.File, error)
	Count(*models.FileFilter) (*int, error)
	Delete(*models.File) error
	Find(id string) (*models.File, error)
}

type Services struct {
	Auth
	Profile
	Users
	Products
	Files
}

func NewServices(repositories *repositories.Repositories) *Services {
	productsService := newProductsService(repositories.Products, repositories.Files)

	rootServices := &Services{
		Auth:     newAuthService(repositories.Users),
		Profile:  newProfileService(repositories.Users),
		Users:    newUsersService(repositories.Users),
		Products: productsService,
		Files:    newFilesService(repositories.Files),
	}

	productsService.rootServices = rootServices

	return rootServices
}
