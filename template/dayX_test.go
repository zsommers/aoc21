package dayX

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = ``

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 0, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 0, B(input))
}
