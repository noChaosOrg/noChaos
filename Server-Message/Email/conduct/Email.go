/**
邮箱操作处理
wsc 2020-8-4
*/
package conduct

import (
	"github.com/go-gomail/gomail"
	"github.com/noChaos1012/noChaos/Message/mail/defs"
)

var m *gomail.Message

/**
发送邮件
*/
func SendMail(mail *defs.Email) error {
	m = gomail.NewMessage()
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)
	m.SetHeader("From", mail.FromEmail)
	m.SetHeader("To", mail.Toers...)
	d := gomail.NewDialer(mail.ServerHost, mail.ServerPort, mail.FromEmail, mail.FromPwd)
	return d.DialAndSend(m)
}
