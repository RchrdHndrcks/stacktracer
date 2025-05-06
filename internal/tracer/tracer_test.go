package tracer_test

import (
	"testing"

	"github.com/RchrdHndrcks/stacktracer"
	"github.com/RchrdHndrcks/stacktracer/internal/tracer"
)

func TestThisCallsOther(t *testing.T) {
	err := tracer.ThisCallsOther()
	if err == nil {
		t.Error("expected error")
	}

	tracedErr := stacktracer.Trace(err)
	if tracedErr == nil {
		t.Error("expected traced error")
	}

	expectedErrMsg := "tracer_test.go:16 - tracer.go:11 - tracer.go:19 - always fails"
	if tracedErr.Error() != expectedErrMsg {
		t.Errorf("expected %q, got %q", expectedErrMsg, tracedErr.Error())
	}
}
