// Package models provides all the structs used in the project
package models

// Response represents the request outcome from the sent email
type Response struct {
	Status  int                 // the status code from request (204,500..)
	Body    string              // more detailed information about the request
	Headers map[string][]string // the readers used in the request
}
