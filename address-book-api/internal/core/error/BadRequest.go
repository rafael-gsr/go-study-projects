package custom_errors

import "errors"

var BadRequestError *GenericError = NewGenericError(
	400,
	errors.New("BadRequestError"),
	"check your requisition, the server cannot handle it",
	map[string]string{},
)
