package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailInfo struct {
	FromMail         string              `json:"fromMail"`
	FromName         string              `json:"fromName"`
	To               string              `json:"to"`
	Subject          string              `json:"subject"`
	PlainTextContent string              `json:"bodyText"`
	HtmlContent      string              `json:"bodyHtml"`
	Identifier       map[string][]string `json:"identifier"`
}

type Response struct {
	Status  int
	Body    string
	Headers map[string][]string
}

func main() {
	lambda.Start(Mail)
}

func Mail(info MailInfo) (Response, error) {
	from := mail.NewEmail(info.FromName, info.FromMail)
	to := mail.NewEmail(info.To, info.To)
	message := mail.NewSingleEmail(from, info.Subject, to, info.PlainTextContent, info.HtmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Sprintf("Failed to send email: %d", info.Identifier)
		log.Println(err)
		return Response{}, err
	}
	return Response{response.StatusCode, response.Body, response.Headers}, err
}
