// Package types contains the main common types of the address book
package types

import custom_errors "address-book-api/internal/core/error"

type IGenericError interface {
	StatusCode() int16
	Err() error
	Message() string
	Extra() map[string]string
	ID() string
	Equals(anotherError *custom_errors.GenericError) bool

	SetExtra(extra map[string]string)
}
