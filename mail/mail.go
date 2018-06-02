// Package mail provides a interface to interact
package mail

import (
	"log"
	"os"

	"github.com/lucasmdrs/go-sendmail/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendMail performs a request to the sendgrid API to send an email with the information from the MailInfo struct passed and returns a Response struct if succeed
func SendMail(info models.MailInfo) (*models.Response, error) {
	from := mail.NewEmail(info.FromName, info.FromMail)
	to := mail.NewEmail(info.To, info.To)
	message := mail.NewSingleEmail(from, info.Subject, to, info.PlainTextContent, info.HTMLContent)

	// Adds any custom arguments to the email information
	// witch can be used to pass an Identifier or some sort of Customer information and
	// it can be retrieved in the sendgrid's console/API later on
	for k, v := range info.CustomArgs {
		message.SetCustomArg(k, v)
	}

	// Start the sendgrid's client to send the email using the API key from the enviroment variable SENDGRID_API_KEY
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Printf("Failed to send email: %d", info.CustomArgs)
		log.Println(err)
		return &models.Response{}, err
	}
	log.Printf("Success to send email: %d", response)
	return &models.Response{response.StatusCode, response.Body, response.Headers}, err
}
