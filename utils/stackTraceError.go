package utils /* Utils */

import (
	"runtime/debug"
)

// # StackTrace is an alias for type returned by debug.Stack()
type StackTrace = []byte

// # StackTraceError is a custom error type with opportunity of saving and passing a stack of calls.
//
// # Description:
//
// Use 'NewSTError(errMes string)' to create new instance.
//
// # Methods:
//
//   - Error() returns string of error message
//   - StackTrace() returns call stack
type StackTraceError struct {
	errMes     string
	stackTrace StackTrace
}

func (ste StackTraceError) Error() string {
	return ste.errMes
}

func (ste StackTraceError) StackTrace() StackTrace {
	return ste.stackTrace
}

func NewSTError(errMes string) *StackTraceError {
	return &StackTraceError{
		errMes:     errMes,
		stackTrace: debug.Stack()}
}
