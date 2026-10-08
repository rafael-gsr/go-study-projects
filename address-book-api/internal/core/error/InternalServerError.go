package error

import "errors"

var InternalServerError *genericError = NewGenericError(
	500,
	errors.New("InternalServerError"),
	"the server cannot handle the requisition",
	map[string]string{},
)
