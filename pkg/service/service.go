package service

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repository"
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

type IUsersService interface {
	CreateUser(*models.User) error
	UpdateUser(*models.User, *models.UserPartial) error
	DeleteUser(*models.User) error
	FindUser(id string) (*models.User, error)
	ListUsers(*repository.PaginationRequest, *repository.Sort, *models.UserFilter) ([]*models.User, error)
	CountUsers(*models.UserFilter) (*int, error)
	ChangeUserPassword(*models.User, string) error
}

type IProductsService interface {
	CreateProduct(*models.Product) error
	UpdateProduct(*models.Product, *models.ProductPartial) error
	DeleteProduct(*models.Product) error
	FindProduct(id string) (*models.Product, error)
	ListProducts(*repository.PaginationRequest, *repository.Sort, *models.ProductFilter) ([]*models.Product, error)
	CountProducts(*models.ProductFilter) (*int, error)
}

type IFilesService interface {
	CreateFile(*models.File) error
	ListFiles(*repository.PaginationRequest, *repository.Sort, *models.FileFilter) ([]*models.File, error)
	CountFiles(*models.FileFilter) (*int, error)
	DeleteFile(*models.File) error
	FindFile(id string) (*models.File, error)
}

type Services struct {
	IAuthService
	IProfileService
	IUsersService
	IProductsService
	IFilesService
}

func NewServices(repositories *repository.Repositories) *Services {
	productsService := newProductsService(repositories)

	RootServices := &Services{
		IAuthService:     newAuthService(repositories),
		IProfileService:  newProfileService(repositories),
		IUsersService:    newUsersService(repositories),
		IProductsService: productsService,
		IFilesService:    newFilesService(repositories),
	}

	productsService.RootServices = RootServices

	return RootServices
}
