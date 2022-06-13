package day17

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `target area: x=20..30, y=-10..-5`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 45, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 112, B(input))
}

func Test_readInput(t *testing.T) {
	want := bounds{
		minX: 20,
		maxX: 30,
		minY: -10,
		maxY: -5,
	}
	assert.Equal(t, want, readInput(input[0]))
}

func Test_x(t *testing.T) {
	tests := []struct {
		initialV int
		time     int
		want     int
	}{
		{6, 1, 6},
		{6, 2, 11},
		{6, 3, 15},
		{6, 4, 18},
		{6, 5, 20},
		{6, 6, 21},
		{6, 7, 21},
		{6, 8, 21},
		{6, 9, 21},
		{6, 10, 21},
		{9, 1, 9},
		{9, 2, 17},
		{9, 3, 24},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d after %d", tt.initialV, tt.time), func(t *testing.T) {
			assert.Equal(t, tt.want, x(tt.initialV, tt.time))
		})
	}
}
