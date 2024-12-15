package services

import (
	"errors"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
)

type ProfileService struct {
	usersRepo repositories.Users
}

func newProfileService(usersRepo repositories.Users) *ProfileService {
	return &ProfileService{usersRepo: usersRepo}
}

func (s *ProfileService) UpdateProfile(u *models.User, data *models.UserPartial) error {
	if *u.Email == adminEmail {
		return errors.New("Administrator can not be updated")
	}

	return s.usersRepo.Update(u, data)
}
