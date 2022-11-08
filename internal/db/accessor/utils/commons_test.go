package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyYear(t *testing.T) {
	assert.NoError(t, VerifyYear("1800"))
	assert.NoError(t, VerifyYear("1895"))
	assert.NoError(t, VerifyYear("1900"))
	assert.NoError(t, VerifyYear("1999"))
	assert.NoError(t, VerifyYear("2000"))
	assert.NoError(t, VerifyYear("2099"))

	assert.Error(t, VerifyYear("1799"))
	assert.Error(t, VerifyYear("2100"))
}
