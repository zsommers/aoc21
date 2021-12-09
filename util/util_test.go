package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Test_readIntString(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{
			s:    "16,1,2,0,4,2,7,1,2,14",
			want: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
		},
		{
			s:    "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			want: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			var result []int
			require.NotPanics(t, func() { result = ReadIntString(tt.s) })
			assert.Equal(t, tt.want, result)
		})
	}
}
