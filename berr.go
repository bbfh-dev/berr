package berr

import (
	"fmt"
	"io"
)

type boxedErr struct {
	label string
	err   error
	ctx   []any
	next  error
}

func New(label string, err error) error {
	if err == nil {
		return nil
	}
	return boxedErr{
		label: label,
		err:   err,
		ctx:   nil,
	}
}

func (boxed boxedErr) HasError() bool {
	if boxed.err == nil {
		return false
	}
	if err, ok := boxed.err.(boxedErr); ok {
		return err.HasError()
	}
	return boxed.err != nil
}

func (boxed boxedErr) Error() string {
	if boxed.err == nil {
		return "<nil>"
	}
	return boxed.Head()
}

func (boxed boxedErr) Head() string {
	return boxed.label + ": " + boxed.err.Error()
}

func (boxed boxedErr) Tail() string {
	if boxed.err == nil {
		return "<nil>"
	}
	if err, ok := boxed.err.(boxedErr); ok {
		return err.Tail()
	}
	return boxed.err.Error()
}
