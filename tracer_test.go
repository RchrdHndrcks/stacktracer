package stacktracer_test

import (
	"errors"
	"testing"

	"github.com/RchrdHndrcks/stacktracer"
)

func thisAlwaysFails() error {
	return errors.New("always fails")
}

type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

func TestTrace(t *testing.T) {
	t.Run("trace", func(t *testing.T) {
		err := thisAlwaysFails()
		if err == nil {
			t.Error("expected error")
		}

		tracedErr := stacktracer.Trace(err)
		if tracedErr == nil {
			t.Error("expected traced error")
		}

		expectedErrMsg := "tracer_test.go:20 - always fails"
		if tracedErr.Error() != expectedErrMsg {
			t.Errorf("expected %q, got %q", expectedErrMsg, tracedErr.Error())
		}
	})

	t.Run("wrapps custom error", func(t *testing.T) {
		err := &customError{message: "always fails"}
		tracedErr := stacktracer.Trace(err)
		if tracedErr == nil {
			t.Error("expected traced error")
		}

		expectedErrMsg := "tracer_test.go:42 - always fails"
		if tracedErr.Error() != expectedErrMsg {
			t.Errorf("expected %q, got %q", expectedErrMsg, tracedErr.Error())
		}

		if !errors.Is(tracedErr, err) {
			t.Error("expected traced error to be the same as the original error")
		}

		var customErr *customError
		if !errors.As(tracedErr, &customErr) {
			t.Error("expected traced error to be the same as the original error")
		}
	})
}
