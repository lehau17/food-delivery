package mailProvider

type emailSender struct {
	From     string
	To       []string
	Password string
	SmtpHost string
	SmtpPort string
}

func NewEmailSender(from string, to []string, password string, smtpPort string, SmtpHost string) *emailSender {
	return &emailSender{From: from, To: to, Password: password, SmtpHost: SmtpHost, SmtpPort: smtpPort}
}

func (s *emailSender) SetTo(to []string) {
	s.To = to
}

type EmailProvider interface {
	SendMail(to []string, templateName, subject string, data interface{}) error
	ParseTemplate(templateFileName string, data interface{}) (string, error)
}
