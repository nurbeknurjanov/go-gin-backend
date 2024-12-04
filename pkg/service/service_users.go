package service

import (
	"errors"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repository"
)

type UsersService struct {
	UsersRepo repository.IUsersRepository
}

func (s *UsersService) CreateUser(u *models.User) error {
	encryptedPassword, err := helpers.EncryptString(*u.Password)
	if err != nil {
		return err
	}

	u.Password = &encryptedPassword
	return s.UsersRepo.CreateUser(u)
}

func (s *UsersService) UpdateUser(u *models.User, data *models.UserPartial) error {
	return s.UsersRepo.UpdateUser(u, data)
}

func (s *UsersService) DeleteUser(u *models.User) error {
	return s.UsersRepo.DeleteUser(u)
}

func (s *UsersService) FindUser(ID string) (*models.User, error) {
	return s.UsersRepo.FindUser(ID)
}

func (s *UsersService) ListUsers(p *repository.PaginationRequest, sort *repository.Sort, f *models.UserFilter) ([]*models.User, error) {
	return s.UsersRepo.ListUsers(p, sort, f)
}
func (s *UsersService) CountUsers(f *models.UserFilter) (*int, error) {
	return s.UsersRepo.CountUsers(f)
}

func (s *UsersService) ChangeUserPassword(u *models.User, password string) error {
	if *u.Email == adminEmail {
		return errors.New("Administrator can not be updated")
	}

	encryptedPassword, err := helpers.EncryptString(password)
	if err != nil {
		return err
	}

	return s.UsersRepo.ChangeUserPassword(u, encryptedPassword)
}

func newUsersService(repositories *repository.Repositories) *UsersService {
	return &UsersService{UsersRepo: repositories}
}
