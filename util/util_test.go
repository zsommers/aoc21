package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckErr(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		shouldPanic bool
	}{
		{
			name:        "No error",
			err:         nil,
			shouldPanic: false,
		},
		{
			name:        "Error",
			err:         errors.New("test"),
			shouldPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panics(t, func() { CheckErr(tt.err) })
			} else {
				assert.NotPanics(t, func() { CheckErr(tt.err) })
			}
		})
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		shouldPanic bool
	}{
		{
			input:    "0",
			expected: 0,
		},
		{
			input:    "10",
			expected: 10,
		},
		{
			input:       "1,000",
			expected:    0,
			shouldPanic: true,
		},
		{
			input:    "-5",
			expected: -5,
		},
		{
			input:       "Not a number",
			expected:    0,
			shouldPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panics(t, func() { Atoi(tt.input) })
			} else {
				var output int
				assert.NotPanics(t, func() { output = Atoi(tt.input) })
				assert.Equal(t, tt.expected, output)
			}
		})
	}
}
