package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestA(t *testing.T) {
	assert.Equal(t, 198, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 230, B(input))
}

func Test_filterList(t *testing.T) {
	type args struct {
		input []string
		index int
		match byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Pass",
			args: args{
				input: []string{"a", "a", "c"},
				index: 0,
				match: 'a',
			},
			want: []string{"a", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, filterList(tt.args.input, tt.args.index, tt.args.match))
		})
	}
}
