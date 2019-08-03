package service

import (
	"../model"
	"fmt"
	"github.com/pkg/errors"
	"net/smtp"
	"strings"
)

const (
	username     = "18861857305@163.com"
	password     = "password"
	host         = "smtp.163.com"
	nickname     = "Aditum"
	user         = "phone@163.com"
	content_type = "Content-Type: text/plain; charset=UTF-8"
)

func SendEmail(email *model.EmailInfo) error {
	title := email.EmailTitle
	content := email.EmailContent
	recipient := email.RecipientAddress
	if title == "" || content == "" || recipient == "" {
		msg := "email is not valid, property may be nil"
		return errors.New(msg)
	}

	auth := smtp.PlainAuth("", username, password, host)
	to := []string{recipient}
	subject := title
	body := content
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.163.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
		return err
	}

	return nil
}
