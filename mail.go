package main

import (
	"net/smtp"
	"strings"

	"github.com/go-gomail/gomail"
	"go.uber.org/zap"
)

type MailMsg struct {
	subject string
	tos     string
	content string
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func (this *MailMsg) send() error {

	d := gomail.Dialer{
		Host: cfg.SMTP.Host,
		Port: cfg.SMTP.Port,
	}
	d.Auth = unencryptedAuth{smtp.PlainAuth("", cfg.SMTP.UserName, cfg.SMTP.Password, cfg.SMTP.Host)}
	d.SSL = false
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.SMTP.UserName)
	m.SetHeader("To", strings.Split(this.tos, ",")...)
	m.SetHeader("Subject", this.subject)
	m.SetBody("text/plain", this.content)
	// Send the email to Bob, Cora and Dan.
	err := d.DialAndSend(m)
	if err != nil {
		logger.Error("send mail fail", zap.String("tos", this.tos), zap.String("subject", this.subject), zap.String("content", this.content), zap.Error(err))
		return err
	}
	logger.Debug("send mail success", zap.String("tos", this.tos), zap.String("subject", this.subject), zap.String("content", this.content))
	return nil

}
