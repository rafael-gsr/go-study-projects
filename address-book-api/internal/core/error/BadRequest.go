package error

import "errors"

var BadRequestError *genericError = NewGenericError(
	400,
	errors.New("BadRequestError"),
	"check your requisition, the server cannot handle it",
	map[string]string{},
)
