package berr_test

import (
	"errors"
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
