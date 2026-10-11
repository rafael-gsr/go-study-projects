// Package types contains the main common types of the address book
package types

import custom_errors "address-book-api/internal/core/error"

type (
	IGenericError interface {
		StatusCode() int16
		Err() error
		Message() string
		Extra() map[string]string
		ID() string
		Equals(anotherError *custom_errors.GenericError) bool

		SetExtra(extra map[string]string)
	}

	Address interface {
		Complement() string
		Block() string
		Street() string
		Number() int16
		City() string
		Country() string

		SetComplement(complement string)
		SetBlock(block string)
		SetStreet(street string)
		SetNumber(number int16)
		SetCity(city string)
		SetCountry(country string)
	}
)
