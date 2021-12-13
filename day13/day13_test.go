package day13

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 17, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 16, B(input))
}

func Test_readInput(t *testing.T) {
	points := map[point]bool{
		{6, 10}:  true,
		{0, 14}:  true,
		{9, 10}:  true,
		{0, 3}:   true,
		{10, 4}:  true,
		{4, 11}:  true,
		{6, 0}:   true,
		{6, 12}:  true,
		{4, 1}:   true,
		{0, 13}:  true,
		{10, 12}: true,
		{3, 4}:   true,
		{3, 0}:   true,
		{8, 4}:   true,
		{1, 10}:  true,
		{2, 14}:  true,
		{8, 10}:  true,
		{9, 0}:   true,
	}
	folds := []fold{
		{"y", 7},
		{"x", 5},
	}
	var p map[point]bool
	var f []fold
	require.NotPanics(t, func() { p, f = readInput(input) })
	assert.Equal(t, points, p)
	assert.ElementsMatch(t, folds, f)
}
