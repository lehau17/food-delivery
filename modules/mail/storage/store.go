package mailstorage

import (
	"context"

	mailProvider "github.com/lehau17/food_delivery/components/mailprovider"
	"github.com/lehau17/food_delivery/components/mailprovider/mail"
)

var (
	mailProviderImple = mail.NewEmailConfig("hau17131203@gmail.com", "quup slyz pzwp ifog", "smtp.gmail.com", "587")
)

type mailStore struct {
	mail mailProvider.EmailProvider
}

func NewMailStore() *mailStore {
	return &mailStore{mail: mailProviderImple}
}

func (s *mailStore) SendMail(ctx context.Context, to []string, templateName string, subject string, data interface{}) error {
	return s.mail.SendMail(to, templateName, subject, data)
}
