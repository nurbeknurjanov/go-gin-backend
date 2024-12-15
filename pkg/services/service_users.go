package services

import (
	"errors"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
)

// agnostic service
// because repo repositories.Users is interface
type UsersService struct {
	repo repositories.Users
}

func newUsersService(repo repositories.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(m *models.User) error {
	encryptedPassword, err := helpers.EncryptString(*m.Password)
	if err != nil {
		return err
	}

	m.Password = &encryptedPassword
	return s.repo.Create(m)
}

func (s *UsersService) Update(m *models.User, data *models.UserPartial) error {
	return s.repo.Update(m, data)
}

func (s *UsersService) Delete(m *models.User) error {
	return s.repo.Delete(m)
}

func (s *UsersService) Find(ID string) (*models.User, error) {
	return s.repo.Find(ID)
}

func (s *UsersService) List(p *repositories.PaginationRequest, sort *repositories.Sort, f *models.UserFilter) ([]*models.User, error) {
	return s.repo.List(p, sort, f)
}
func (s *UsersService) Count(f *models.UserFilter) (*int, error) {
	return s.repo.Count(f)
}

func (s *UsersService) ChangeUserPassword(m *models.User, password string) error {
	if *m.Email == adminEmail {
		return errors.New("Administrator can not be updated")
	}

	encryptedPassword, err := helpers.EncryptString(password)
	if err != nil {
		return err
	}

	return s.repo.ChangeUserPassword(m, encryptedPassword)
}
