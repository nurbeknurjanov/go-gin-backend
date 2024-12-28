package services

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/kafka"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"strings"
)

var (
	templateRegistration = "Welcome {name}"
	topicRegistration    = "topic-registration"
)

type MailingService struct {
	producer *kafka.Producer
}

func newMailingService(producer *kafka.Producer) *MailingService {
	return &MailingService{producer}
}

func (m *MailingService) SendRegistrationMessage(u *models.User) error {
	return nil
	messageRegistration := strings.Replace(templateRegistration, "{name}", *u.Name, -1)
	return m.producer.Produce(messageRegistration, topicRegistration, nil)
}
