package otp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOtp(t *testing.T) {
	otp := generateOtp()
	assert.NotNil(t, otp)
	assert.NotEmpty(t, otp)
}
