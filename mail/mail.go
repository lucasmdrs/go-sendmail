package mail

import (
	"log"
	"os"

	"github.com/lucasmdrs/go-sendmail/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(info models.MailInfo) (*models.Response, error) {
	from := mail.NewEmail(info.FromName, info.FromMail)
	to := mail.NewEmail(info.To, info.To)
	message := mail.NewSingleEmail(from, info.Subject, to, info.PlainTextContent, info.HTMLContent)
	for k, v := range info.CustomArgs {
		message.SetCustomArg(k, v)
	}
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Printf("Failed to send email: %s", info.CustomArgs)
		log.Println(err)
		return &models.Response{}, err
	}
	log.Printf("Success to send email: %+v", response)
	return &models.Response{response.StatusCode, response.Body, response.Headers}, err
}
