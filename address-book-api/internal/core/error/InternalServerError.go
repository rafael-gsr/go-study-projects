package custom_errors

import "errors"

var InternalServerError *GenericError = NewGenericError(
	500,
	errors.New("InternalServerError"),
	"the server cannot handle the requisition",
	map[string]string{},
)
