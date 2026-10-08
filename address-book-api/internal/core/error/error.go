// Package error contains the possible errors in the application and related resources
package error

import (
	"fmt"
	"strings"
)

type genericError struct {
	id         string
	statusCode int16
	err        error
	message    string
	extra      map[string]string
}

func (g *genericError) StatusCode() int16                      { return g.statusCode }
func (g *genericError) Err() error                             { return g.err }
func (g *genericError) Message() string                        { return g.message }
func (g *genericError) Extra() map[string]string               { return g.extra }
func (g *genericError) ID() string                             { return g.id }
func (g *genericError) Equals(anotherError *genericError) bool { return g.id == anotherError.ID() }

func (g *genericError) SetExtra(extra map[string]string) { g.extra = extra }

func NewGenericError(
	statusCode int16,
	err error,
	message string,
	extra map[string]string,
) *genericError {
	return &genericError{
		fmt.Sprintf("%d-%s-%s", statusCode, strings.TrimSpace(message), strings.TrimSpace(err.Error())),
		statusCode,
		err,
		message,
		extra,
	}
}
