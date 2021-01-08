package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistingFile(t *testing.T) {
	e := Exists("./testdata/existing_file")
	assert.True(t, e)
}

func TestNonExistingFile(t *testing.T) {
	e := Exists("./testdata/non_existing_file")
	assert.False(t, e)
}
