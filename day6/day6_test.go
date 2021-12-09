package day6

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `3,4,3,1,2`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 5934, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 26984457539, B(input))
}

func Test_update(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{0, 1, 1, 2, 1, 0, 0, 0, 0},
			want:  []int{1, 1, 2, 1, 0, 0, 0, 0, 0},
		},
		{
			input: []int{1, 1, 2, 1, 0, 0, 0, 0, 0},
			want:  []int{1, 2, 1, 0, 0, 0, 1, 0, 1},
		},
		{
			input: []int{1, 2, 1, 0, 0, 0, 1, 0, 1},
			want:  []int{2, 1, 0, 0, 0, 1, 1, 1, 1},
		},
		{
			input: []int{2, 1, 0, 0, 0, 1, 1, 1, 1},
			want:  []int{1, 0, 0, 0, 1, 1, 3, 1, 2},
		},
		{
			input: []int{1, 0, 0, 0, 1, 1, 3, 1, 2},
			want:  []int{0, 0, 0, 1, 1, 3, 2, 2, 1},
		},
		{
			input: []int{0, 0, 0, 1, 1, 3, 2, 2, 1},
			want:  []int{0, 0, 1, 1, 3, 2, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.want, update(tt.input))
		})
	}
}

/*
func Test_update(t *testing.T) {
	tests := []struct {
		state int
		days  int
		want  int
	}{
		{
			state: 2,
			days:  1,
			want:  1,
		},
		{
			state: 2,
			days:  2,
			want:  1,
		},
		{
			state: 2,
			days:  3,
			want:  2,
		},
		{
			state: 2,
			days:  9,
			want:  2,
		},
		{
			state: 2,
			days:  10,
			want:  3,
		},
		{
			state: 2,
			days:  11,
			want:  3,
		},
		{
			state: 2,
			days:  12,
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d day %d", tt.state, tt.days), func(t *testing.T) {
			assert.Equal(t, tt.want, update(tt.state, tt.days))
		})
	}
}*/

func Test_readInput(t *testing.T) {
	//assert.Equal(t, []int{3, 4, 3, 1, 2}, readInput(input[0]))
	assert.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, readInput(input[0]))
}
