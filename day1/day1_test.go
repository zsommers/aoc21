package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"199",
	"200",
	"208",
	"210",
	"200",
	"207",
	"240",
	"269",
	"260",
	"263",
}

func TestDay1A(t *testing.T) {
	assert.Equal(t, 7, A(input))
}

func TestDay1B(t *testing.T) {
	assert.Equal(t, 5, B(input))
}
