package tracer

import (
	"errors"

	"github.com/RchrdHndrcks/stacktracer"
)

func ThisCallsOther() error {
	if err := Other(); err != nil {
		return stacktracer.Trace(err)
	}

	return nil
}

func Other() error {
	if err := Another(); err != nil {
		return stacktracer.Trace(err)
	}

	return nil
}

func Another() error {
	return errors.New("always fails")
}
