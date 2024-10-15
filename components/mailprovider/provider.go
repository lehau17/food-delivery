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

type DataSendMail struct {
	To           []string
	Subject      string
	TemplateName string
	Payload      interface{}
}

func (d *DataSendMail) GetPayLoadSendEmail() interface{} {
	return d.Payload
}

func (d *DataSendMail) GetTo() []string {
	return d.To
}

func (d *DataSendMail) GetSubject() string {
	return d.Subject
}

func (d *DataSendMail) GetTemplateName() string {
	return d.TemplateName
}
