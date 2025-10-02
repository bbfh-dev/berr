package berr_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/bbfh-dev/berr"
	"gotest.tools/assert"
)

func TestBerrNesting(test *testing.T) {
	assert.DeepEqual(test, berr.New("example", nil), nil)
	assert.DeepEqual(test, berr.New("example", errors.New("test")).Error(), "example: test")
	assert.DeepEqual(test, berr.New("example", berr.New("nested", errors.New("test"))).Error(), "example: nested: test")
	assert.DeepEqual(test, berr.New("example", berr.New("nested", nil)), nil)
}

func TestBerr(test *testing.T) {
	err := berr.WithContext(
		"another example",
		berr.WithContext(
			"Hello World!",
			errors.New("Yet another error!"),
			"c",
			"Something Something",
			"d",
			map[string]bool{"x": true, "y": false, "z": true},
		),
		"a",
		123,
		"b",
		456,
	)
	fmt.Println("```test")
	berr.Fexpand(os.Stdout, err)
	fmt.Println("\n```")
}

func TestIgnore(test *testing.T) {
	var buffer bytes.Buffer
	_, err := buffer.ReadByte()
	assert.Equal(test, berr.New("shouldn't fail", err).Ignore(io.EOF).HasError(), false)
}
