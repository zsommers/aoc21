package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input2 = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestDay2A(t *testing.T) {
	assert.Equal(t, 150, A(input2))
}

func TestDay2B(t *testing.T) {
	assert.Equal(t, 900, B(input2))
}
