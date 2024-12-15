package service

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
)

const (
	adminEmail = "nurbek.nurjanov@mail.ru"
)

type IAuthService interface {
	Login(string, string) (*models.Tokens, error)
	GetAccessToken(*models.User) (string, error)
}

type IProfileService interface {
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

type IProductsService interface {
	CreateProduct(*models.Product) error
	UpdateProduct(*models.Product, *models.ProductPartial) error
	DeleteProduct(*models.Product) error
	FindProduct(id string) (*models.Product, error)
	ListProducts(*repositories.PaginationRequest, *repositories.Sort, *models.ProductFilter) ([]*models.Product, error)
	CountProducts(*models.ProductFilter) (*int, error)
}

type IFilesService interface {
	CreateFile(*models.File) error
	ListFiles(*repositories.PaginationRequest, *repositories.Sort, *models.FileFilter) ([]*models.File, error)
	CountFiles(*models.FileFilter) (*int, error)
	DeleteFile(*models.File) error
	FindFile(id string) (*models.File, error)
}

type Services struct {
	IAuthService
	IProfileService
	Users
	IProductsService
	IFilesService
}

func NewServices(repositories *repositories.Repositories) *Services {
	productsService := newProductsService(repositories)

	RootServices := &Services{
		IAuthService:     newAuthService(repositories),
		IProfileService:  newProfileService(repositories),
		Users:            newUsersSqlService(repositories.Users),
		IProductsService: productsService,
		IFilesService:    newFilesService(repositories),
	}

	productsService.RootServices = RootServices

	return RootServices
}
