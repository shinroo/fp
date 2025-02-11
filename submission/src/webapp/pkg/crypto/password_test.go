package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		expectError bool
	}{
		{
			name:        "valid password",
			password:    "mysecretpassword",
			expectError: false,
		},
		{
			name:        "empty password",
			password:    "",
			expectError: false, // bcrypt allows empty passwords
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := HashPassword(tt.password)

			if tt.expectError {
				require.Error(t, err, "HashPassword() should return an error")
			} else {
				require.NoError(t, err, "HashPassword() should not return an error")
				assert.NotEmpty(t, hashedPassword, "HashPassword() should not return an empty string")
				assert.NotEqual(t, tt.password, hashedPassword, "HashPassword() should not return the same string as the input password")
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	validPassword := "mysecretpassword"
	hashedPassword, err := HashPassword(validPassword)
	require.NoError(t, err, "HashPassword() failed during setup")

	tests := []struct {
		name        string
		hashedPass  string
		inputPass   string
		expectError bool
	}{
		{
			name:        "correct password",
			hashedPass:  hashedPassword,
			inputPass:   validPassword,
			expectError: false,
		},
		{
			name:        "incorrect password",
			hashedPass:  hashedPassword,
			inputPass:   "wrongpassword",
			expectError: true,
		},
		{
			name:        "empty hashed password",
			hashedPass:  "",
			inputPass:   validPassword,
			expectError: true,
		},
		{
			name:        "empty input password",
			hashedPass:  hashedPassword,
			inputPass:   "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ComparePassword(tt.hashedPass, tt.inputPass)

			if tt.expectError {
				assert.Error(t, err, "ComparePassword() should return an error")
			} else {
				assert.NoError(t, err, "ComparePassword() should not return an error")
			}
		})
	}
}
