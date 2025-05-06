package stacktracer_test

import (
	"errors"
	"testing"

	"github.com/RchrdHndrcks/stacktracer"
)

func thisAlwaysFails() error {
	return errors.New("always fails")
}

func TestTrace(t *testing.T) {
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
}
