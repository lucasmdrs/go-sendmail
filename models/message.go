package models

type MailInfo struct {
	FromMail         string            `json:"fromMail"`
	FromName         string            `json:"fromName"`
	To               string            `json:"to"`
	Subject          string            `json:"subject"`
	PlainTextContent string            `json:"bodyText"`
	HtmlContent      string            `json:"bodyHtml"`
	CustomArgs       map[string]string `json:"customArgs"`
}
