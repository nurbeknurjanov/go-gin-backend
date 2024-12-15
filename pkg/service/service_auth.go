package service

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	usersRepo repositories.IUsersRepository
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

func newAuthService(repositories *repositories.Repositories) *AuthService {
	return &AuthService{usersRepo: repositories}
}
