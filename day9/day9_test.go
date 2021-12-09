package day9

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `2199943210
3987894921
9856789892
8767896789
9899965678`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 15, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 1134, B(input))
}

func Test_readInput(t *testing.T) {
	var result [][]int
	expected := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	require.NotPanics(t, func() { result = readInput(input) })
	assert.Equal(t, expected, result)
}

func Test_checkLowPoint(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want bool
	}{
		{
			x:    1,
			y:    0,
			want: true,
		},
		{
			x:    9,
			y:    0,
			want: true,
		},
		{
			x:    2,
			y:    2,
			want: true,
		},
		{
			x:    6,
			y:    4,
			want: true,
		},
		{
			x:    0,
			y:    1,
			want: false,
		},
		{
			x:    4,
			y:    0,
			want: false,
		},
		{
			x:    8,
			y:    2,
			want: false,
		},
		{
			x:    3,
			y:    4,
			want: false,
		},
	}
	m := readInput(input)
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.x, tt.y), func(t *testing.T) {
			var result bool
			require.NotPanics(t, func() { result = checkLowPoint(m, tt.x, tt.y) })
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_findLowPoints(t *testing.T) {
	m := readInput(input)
	want := []point{
		{1, 0},
		{9, 0},
		{2, 2},
		{6, 4},
	}
	var result []point
	require.NotPanics(t, func() { result = findLowPoints(m) })
	assert.ElementsMatch(t, want, result)
}

func Test_sizeBasin(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{
			x:    1,
			y:    0,
			want: 3,
		},
		{
			x:    9,
			y:    0,
			want: 9,
		},
		{
			x:    2,
			y:    2,
			want: 14,
		},
		{
			x:    6,
			y:    4,
			want: 9,
		},
	}
	m := readInput(input)
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.x, tt.y), func(t *testing.T) {
			var result int
			require.NotPanics(t, func() { result, _ = sizeBasin(m, tt.x, tt.y, []point{}) })
			assert.Equal(t, tt.want, result)
		})
	}
}
