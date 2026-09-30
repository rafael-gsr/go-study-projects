// Package service contains the functions to handle the internal flows
package service

type AddressBookService struct{}

func (s *AddressBookService) Pong() map[string]string {
	response := map[string]string{
		"message": "pong",
	}

	return response
}
