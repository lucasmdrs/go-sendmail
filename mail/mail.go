// Package mail provides a interface to interact
package mail

import (
	"net/http"
	"os"

	"github.com/lucasmdrs/go-sendmail/logger"
	"github.com/lucasmdrs/go-sendmail/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendMail performs a request to the sendgrid API to send an email with the information from the MailInfo struct passed and returns a Response struct if succeed
func SendMail(info models.MailInfo) (*models.Response, error) {
	log := logger.DefaultLogger()

	log.Debug("Setting Sender:", info.FromMail, info.FromName)
	from := mail.NewEmail(info.FromName, info.FromMail)
	log.Debug("Setting Receiver:", info.To)
	to := mail.NewEmail(info.To, info.To)
	log.Debug("Creating Message:", info.Subject, info.PlainTextContent, info.HTMLContent)
	message := mail.NewSingleEmail(from, info.Subject, to, info.PlainTextContent, info.HTMLContent)

	// Adds any custom arguments to the email information
	// witch can be used to pass an Identifier or some sort of Customer information and
	// it can be retrieved in the sendgrid's console/API later on
	for k, v := range info.CustomArgs {
		log.Debug("Adding Custom Args:", k, v)
		message.SetCustomArg(k, v)
	}

	// Start the sendgrid's client to send the email using the API key from the enviroment variable SENDGRID_API_KEY
	log.Debug("Starting Sendgrid Client")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	log.Debug("Sending Email")
	response, err := client.Send(message)
	if err != nil {
		log.Error("Failed to send email: %s", err.Error())
		log.Debug(message)
		return &models.Response{}, err
	}
	if response.StatusCode > http.StatusAccepted {
		log.Error("Failed to send email:", response)
	}
	log.Info("Successfully sent email")
	return &models.Response{response.StatusCode, response.Body, response.Headers}, err
}
