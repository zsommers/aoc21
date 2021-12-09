package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 5, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 12, B(input))
}

func Test_createRange(t *testing.T) {
	tests := []struct {
		name string
		a    point
		b    point
		want []point
	}{
		{
			name: "horizontal",
			a:    point{1, 4},
			b:    point{4, 4},
			want: []point{
				{1, 4},
				{2, 4},
				{3, 4},
				{4, 4},
			},
		},
		{
			name: "vertical",
			a:    point{4, 1},
			b:    point{4, 4},
			want: []point{
				{4, 1},
				{4, 2},
				{4, 3},
				{4, 4},
			},
		},
		{
			name: "diagonal up",
			a:    point{1, 1},
			b:    point{4, 4},
			want: []point{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
			},
		},
		{
			name: "diagonal down",
			a:    point{4, 1},
			b:    point{1, 4},
			want: []point{
				{4, 1},
				{3, 2},
				{2, 3},
				{1, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, createRange(tt.a, tt.b))
		})
	}
}

func Test_readSegment(t *testing.T) {
	tests := []struct {
		s     string
		wantA point
		wantB point
	}{
		{
			s:     "0,9 -> 5,9",
			wantA: point{0, 9},
			wantB: point{5, 9},
		},
		{
			s:     "7,0 -> 7,4",
			wantA: point{7, 0},
			wantB: point{7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			var a, b point
			require.NotPanics(t, func() { a, b = readSegment(tt.s) })
			assert.Equal(t, tt.wantA, a)
			assert.Equal(t, tt.wantB, b)
		})
	}
}
