package email

import (
	gomail "gopkg.in/gomail.v2"
	not "opensoach.com/engine/notification"
)

func SendEmail(iEmail not.INotificationEmail) error {

	emailOpt := iEmail.GetEmailOptions()
	//	emailOpt.From = "support@opensoach.com"
	//	emailOpt.To = []string{"mandar.bhong@opensoach.com"}
	//	emailOpt.Subject = "Hello!"
	//	emailOpt.Body = "Hello <b>Bob</b> and <i>Cora</i>!"

	//	emailOpt.SMTPAddress = "send.one.com"
	//	emailOpt.SMTPPort = 25
	//	emailOpt.SMTPUsername = "support@opensoach.com"
	//	emailOpt.SMTPPassword = "opensoach.support@123"

	m := gomail.NewMessage()
	m.SetHeader("From", emailOpt.From)
	m.SetHeader("To", emailOpt.To...)
	m.SetHeader("Subject", emailOpt.Subject)
	m.SetBody("text/html", emailOpt.Body)

	for _, attachmentPath := range emailOpt.Attachment {
		m.Attach(attachmentPath)
	}

	d := gomail.NewDialer(emailOpt.SMTPAddress, emailOpt.SMTPPort, emailOpt.SMTPUsername, emailOpt.SMTPPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
