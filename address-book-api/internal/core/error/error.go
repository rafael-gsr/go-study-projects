// Package error contains the possible errors in the application and related resources
package custom_errors

import (
	"fmt"
	"strings"
)

type GenericError struct {
	id         string
	statusCode int16
	err        error
	message    string
	extra      map[string]string
}

func (g *GenericError) StatusCode() int16                      { return g.statusCode }
func (g *GenericError) Err() error                             { return g.err }
func (g *GenericError) Message() string                        { return g.message }
func (g *GenericError) Extra() map[string]string               { return g.extra }
func (g *GenericError) ID() string                             { return g.id }
func (g *GenericError) Equals(anotherError *GenericError) bool { return g.id == anotherError.ID() }

func (g *GenericError) SetExtra(extra map[string]string) { g.extra = extra }

func NewGenericError(
	statusCode int16,
	err error,
	message string,
	extra map[string]string,
) *GenericError {
	return &GenericError{
		fmt.Sprintf("%d-%s-%s", statusCode, strings.TrimSpace(message), strings.TrimSpace(err.Error())),
		statusCode,
		err,
		message,
		extra,
	}
}
