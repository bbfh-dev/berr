package berr

import (
	"io"
	"strings"
)

func WithContext(label string, err error, args ...any) error {
	if err == nil {
		return nil
	}

	boxed := New(label, err).(boxedErr)
	boxed.ctx = args
	return boxed
}

func Expand(err error) string {
	var builder strings.Builder
	Fexpand(&builder, err)
	return builder.String()
}

func Fexpand(writer io.Writer, err error) {
	if err == nil {
		writer.Write([]byte("<nil>"))
		return
	}

	writer.Write([]byte("[Error]\n"))
	writer.Write([]byte(err.Error()))

	if boxed, ok := err.(boxedErr); ok {
		writer.Write([]byte("\n\n[Trace]\n"))
		boxed.Expand(writer, 1)
	}
}
