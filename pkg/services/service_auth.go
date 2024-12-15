package services

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	usersRepo repositories.Users
}

func newAuthService(usersRepo repositories.Users) *AuthService {
	return &AuthService{usersRepo: usersRepo}
}

func (s *AuthService) Login(email string, password string) (*models.Tokens, error) {
	u, err := s.usersRepo.FindByEmail(email)
	if err != nil {
		logrus.Info(err)
		return nil, errLogin
	}

	if err := u.ValidatePassword(password); err != nil {
		logrus.Info(err)
		return nil, errLogin
	}

	u.Password = nil

	return models.GenerateTokens(u), nil
}

func (*AuthService) GetAccessToken(u *models.User) (string, error) {
	return models.GenerateAccessToken(u), nil
}
