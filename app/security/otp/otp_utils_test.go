package otp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOtp(t *testing.T) {
	otp := GenerateOtp()
	assert.NotNil(t, otp)
	assert.NotEmpty(t, otp)
}
