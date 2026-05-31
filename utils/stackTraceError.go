package utils /* Utils */

import (
	"runtime/debug"
)

// TODO docs
type StackTrace = []byte

// TODO docs
type StackTraceError struct {
	errMes     string
	stackTrace StackTrace
}

// TODO docs
func (ste StackTraceError) Error() string {
	return ste.errMes
}

// TODO docs
func (ste StackTraceError) StackTrace() StackTrace {
	return ste.stackTrace
}

// TODO docs
func NewSTError(errMes string) *StackTraceError {
	return &StackTraceError{
		errMes:     errMes,
		stackTrace: debug.Stack()}
}
