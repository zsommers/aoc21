package day7

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zsommers/aoc21/util"
)

var rawInput = `16,1,2,0,4,2,7,1,2,14`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 37, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 168, B(input))
}

func Test_costA(t *testing.T) {
	tests := []struct {
		d    int
		want int
	}{
		{
			d:    1,
			want: 41,
		},
		{
			d:    2,
			want: 37,
		},
		{
			d:    3,
			want: 39,
		},
		{
			d:    10,
			want: 71,
		},
	}
	ps := util.ReadIntString(input[0])
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.d), func(t *testing.T) {
			var result int
			require.NotPanics(t, func() { result = costA(ps, tt.d) })
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_costB(t *testing.T) {
	tests := []struct {
		d    int
		want int
	}{
		{
			d:    2,
			want: 206,
		},
		{
			d:    5,
			want: 168,
		},
	}
	ps := util.ReadIntString(input[0])
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.d), func(t *testing.T) {
			var result int
			require.NotPanics(t, func() { result = costB(ps, tt.d) })
			assert.Equal(t, tt.want, result)
		})
	}
}
