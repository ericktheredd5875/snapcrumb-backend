package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindProjectRoot(t *testing.T) {
	root, err := FindProjectRoot("")
	assert.NoError(t, err)
	assert.Equal(t, "C:/CodeBases/snapcrumb-backend", root)
}
