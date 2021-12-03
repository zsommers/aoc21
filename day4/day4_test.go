package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{}

func TestA(t *testing.T) {
	assert.Equal(t, 198, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 230, B(input))
}
