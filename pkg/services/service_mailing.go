package services

import (
	"fmt"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
)

type MailingService struct {
}

func newMailingService() *MailingService {
	return &MailingService{}
}

func (m *MailingService) SendRegistrationMessage(u *models.User) error {
	fmt.Println("Send")
	return nil
}
