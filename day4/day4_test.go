package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

var input = strings.Split(rawInput, "\n")

func flattenBoard(b board) []boardSpace {
	bs := []boardSpace{}
	for _, r := range b {
		bs = append(bs, r...)
	}
	return bs
}

func TestA(t *testing.T) {
	assert.Equal(t, 4512, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 1924, B(input))
}

func Test_readBoards(t *testing.T) {
	expected := []board{
		{
			{{22, false}, {13, false}, {17, false}, {11, false}, {0, false}},
			{{8, false}, {2, false}, {23, false}, {4, false}, {24, false}},
			{{21, false}, {9, false}, {14, false}, {16, false}, {7, false}},
			{{6, false}, {10, false}, {3, false}, {18, false}, {5, false}},
			{{1, false}, {12, false}, {20, false}, {15, false}, {19, false}},
		},
		{
			{{3, false}, {15, false}, {0, false}, {2, false}, {22, false}},
			{{9, false}, {18, false}, {13, false}, {17, false}, {5, false}},
			{{19, false}, {8, false}, {7, false}, {25, false}, {23, false}},
			{{20, false}, {11, false}, {10, false}, {24, false}, {4, false}},
			{{14, false}, {21, false}, {16, false}, {12, false}, {6, false}},
		},
		{
			{{14, false}, {21, false}, {17, false}, {24, false}, {4, false}},
			{{10, false}, {16, false}, {15, false}, {9, false}, {19, false}},
			{{18, false}, {8, false}, {23, false}, {26, false}, {20, false}},
			{{22, false}, {11, false}, {13, false}, {6, false}, {5, false}},
			{{2, false}, {0, false}, {12, false}, {3, false}, {7, false}},
		},
	}

	var result []board
	require.NotPanics(t, func() { result = readBoards(input) })
	require.Equal(t, len(expected), len(result))
	assert.Equal(t, expected, result)
}

func Test_board_mark(t *testing.T) {
	input := board{
		{{22, false}, {13, false}, {17, false}, {11, false}, {0, false}},
		{{8, false}, {2, false}, {23, false}, {4, false}, {24, false}},
		{{21, false}, {9, false}, {14, false}, {16, false}, {7, false}},
		{{6, false}, {10, false}, {3, false}, {18, false}, {5, false}},
		{{1, false}, {12, false}, {20, false}, {15, false}, {19, false}},
	}
	tests := []struct {
		name string
		b    board
		draw int
		want bool
	}{
		{
			name: "space in board",
			b:    input,
			draw: 14,
			want: true,
		},
		{
			name: "space not in board",
			b:    input,
			draw: 77,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result bool
			require.NotPanics(t, func() { result = tt.b.mark(tt.draw) })
			assert.Equal(t, tt.want, result)
			if tt.want {
				assert.Contains(t, flattenBoard(tt.b), boardSpace{tt.draw, true})
			} else {
				assert.NotContains(t, flattenBoard(tt.b), boardSpace{tt.draw, false})
				assert.NotContains(t, flattenBoard(tt.b), boardSpace{tt.draw, true})
			}
		})
	}
}

func Test_board_checkWin(t *testing.T) {
	tests := []struct {
		name string
		b    board
		want bool
	}{
		{
			name: "no winner",
			b: board{
				{{22, false}, {13, true}, {17, false}, {11, false}, {0, false}},
				{{8, true}, {2, false}, {23, false}, {4, false}, {24, false}},
				{{21, false}, {9, false}, {14, true}, {16, true}, {7, false}},
				{{6, false}, {10, false}, {3, false}, {18, true}, {5, false}},
				{{1, false}, {12, false}, {20, false}, {15, true}, {19, false}},
			},
			want: false,
		},
		{
			name: "row winner",
			b: board{
				{{22, false}, {13, true}, {17, false}, {11, false}, {0, false}},
				{{8, true}, {2, false}, {23, false}, {4, false}, {24, false}},
				{{21, true}, {9, true}, {14, true}, {16, true}, {7, true}},
				{{6, false}, {10, false}, {3, false}, {18, true}, {5, false}},
				{{1, false}, {12, false}, {20, false}, {15, true}, {19, false}},
			},
			want: true,
		},
		{
			name: "column winner",
			b: board{
				{{22, false}, {13, true}, {17, false}, {11, true}, {0, false}},
				{{8, true}, {2, false}, {23, false}, {4, true}, {24, false}},
				{{21, false}, {9, false}, {14, true}, {16, true}, {7, false}},
				{{6, false}, {10, false}, {3, false}, {18, true}, {5, false}},
				{{1, false}, {12, false}, {20, false}, {15, true}, {19, false}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.b.checkWin())
		})
	}
}
