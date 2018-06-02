// Package models provides all the structs used in the project
package models

// MailInfo represents the email information
type MailInfo struct {
	FromMail         string            `json:"fromMail"`   // sender's email address
	FromName         string            `json:"fromName"`   // sender's name
	To               string            `json:"to"`         // receiver's email address
	Subject          string            `json:"subject"`    // email subject
	PlainTextContent string            `json:"bodyText"`   // plain text content for the email body
	HTMLContent      string            `json:"bodyHtml"`   // html content for the email body
	CustomArgs       map[string]string `json:"customArgs"` // any argument to be included in the email ( metadata )
}
