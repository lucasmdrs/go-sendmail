package models

type Response struct {
	Status  int
	Body    string
	Headers map[string][]string
}
