package service

import (
	"errors"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
)

type ProfileService struct {
	usersRepo repositories.IUsersRepository
}

func (s *ProfileService) UpdateProfile(u *models.User, data *models.UserPartial) error {
	if *u.Email == adminEmail {
		return errors.New("Administrator can not be updated")
	}

	return s.usersRepo.UpdateUser(u, data)
}

func newProfileService(repositories *repositories.Repositories) *ProfileService {
	return &ProfileService{usersRepo: repositories}
}
