package stacktracer

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

// Trace wraps the passed error with a stack trace.
//
// It returns a new error with the stack trace added.
// Example:
//
//	err := errors.New("always fails")
//	if err != nil {
//		return stacktracer.Trace(err)
//	}
//
// The error message will be something like:
//
//	file.go:line - always fails
func Trace(err error) error {
	return errors.New(trace(err.Error()))
}

// trace adds the file and line number to the passed message.
func trace(msg string) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return msg
	}
	filename := filepath.Base(file)
	return fmt.Sprintf("%s:%d - %s", filename, line, msg)
}
