package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	mail "github.com/lucasmdrs/go-sendmail/mail"
)

func main() {
	lambda.Start(mail.SendMail)
}
