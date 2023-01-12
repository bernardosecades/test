//go:build unit
// +build unit

package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidUUID(t *testing.T) {
	assert.True(t, IsValidUUID("b26e5791-39c2-402a-b7ca-59105dc474e6"))
}


func TestIsInvalidUUID(t *testing.T) {
	assert.False(t, IsValidUUID("b26e5791-39c2-402a-b7ca"))
	assert.False(t, IsValidUUID(""))
	assert.False(t, IsValidUUID("lololo"))
}