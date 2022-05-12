package common

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	logger "github.com/Almazatun/golephant/pkg/logger"
)

var (
	// Sender data.
	from     = os.Getenv("SMTP_MAIL_FROM")
	password = os.Getenv("SMTP_MAIL_PASSWORD")
	// Smtp server configuration.
	smtpHost = os.Getenv("SMTP_MAIL_HOST")
	smtpPort = os.Getenv("SMTP_MAIL_PORT")

	frontendBaseUrl = os.Getenv("FRONTEND_BASE_URL")
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func SendEmail(to, ressetPasswordToken string) {
	// Receiver
	var reveiver []string
	reveiver = append(reveiver, to)

	subject := "Simple HTML mail"
	body := `<a><b>` + frontendBaseUrl + `\` + ressetPasswordToken + `</b></p>`

	// Request
	request := Mail{
		Sender:  from,
		To:      reveiver,
		Subject: subject,
		Body:    body,
	}

	msg := BuildMessage(request)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, reveiver, []byte(msg))
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info("Email Sent Successfully" + to)
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
