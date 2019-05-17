package notification

type EmailOptions struct {
	To           []string
	From         string
	CC           []string
	Subject      string
	Body         string
	Attachment   []string
	SMTPAddress  string
	SMTPUsername string
	SMTPPassword string
	SMTPPort     int
}
