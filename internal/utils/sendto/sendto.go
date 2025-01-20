package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/twinbeard/goLearning/global"
	"go.uber.org/zap"
)

// This is the SMTP server configuration for sending emails => Lấy trên google nhé
const (
	SMTPHost     = "smtp.gmail.com" // SMTP default server
	SMTPPort     = "25"
	SMTPUsername = "tlttmtf@gmail.com"
	SMTPPassword = "qdhd kash jseh fkwv"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	// Build the email message
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

// Send email in text format => Gửi email với nội dung text thông thường - 0 dùng
func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp), // Nội dung email là OPT dưới dạng text
	}
	// Build the email message
	messageMail := BuildMessage(contentEmail)
	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	//  Send the email to the SMTP server in port 587
	err := smtp.SendMail(SMTPHost+":587", auth, from, to, []byte(messageMail))
	// Check for errors
	if err != nil {
		global.Logger.Error("Email send failed.", zap.Error(err))
		return err
	}
	return nil
}

// C2: Send email in HTML format => Recommend vì nó đẹp và nhanh

func SendTemplateEmailOtp(
	to []string,
	from string,
	nameTemplate string,
	dataTemplate map[string]interface{}) error {
	// Create template email
	htmlBody, err := getEmailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}
	// Send email with html template
	return send(to, from, htmlBody)
}

func getEmailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)                                                            // Tạo một buffer để lưu nội dung email
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate)) // Parse file template at path "templates-email/nameTemplate"

	err := t.Execute(htmlTemplate, dataTemplate) // Execute template với data truyền vào
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate, // Nội dung email là html template
	}
	// Build the email message
	messageMail := BuildMessage(contentEmail)
	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	//  Send the email to the SMTP server in port 587
	err := smtp.SendMail(SMTPHost+":587", auth, from, to, []byte(messageMail))
	// Check for errors
	if err != nil {
		global.Logger.Error("Email send failed.", zap.Error(err))
		return err
	}
	return nil
}
