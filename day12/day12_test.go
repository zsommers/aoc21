package day12

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 226, A(input))
}

func TestB(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "baby",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			want: 36,
		},
		{
			name: "momma",
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			want: 103,
		},
		{
			name:  "daddy",
			input: input,
			want:  3509,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int
			require.NotPanics(t, func() { result = B(tt.input) })
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_readInput(t *testing.T) {
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
	want := map[string][]string{
		"start": {"A", "b"},
		"A":     {"start", "c", "b", "end"},
		"b":     {"start", "A", "d", "end"},
		"c":     {"A"},
		"d":     {"b"},
		"end":   {"A", "b"},
	}
	assert.Equal(t, want, readInput(input))
}

func Test_path(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name: "baby",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			want: []string{
				"start,A,b,A,c,A,end",
				"start,A,b,A,end",
				"start,A,b,end",
				"start,A,c,A,b,A,end",
				"start,A,c,A,b,end",
				"start,A,c,A,end",
				"start,A,end",
				"start,b,A,c,A,end",
				"start,b,A,end",
				"start,b,end",
			},
		},
		{
			name: "momma",
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			want: []string{
				"start,HN,dc,HN,end",
				"start,HN,dc,HN,kj,HN,end",
				"start,HN,dc,end",
				"start,HN,dc,kj,HN,end",
				"start,HN,end",
				"start,HN,kj,HN,dc,HN,end",
				"start,HN,kj,HN,dc,end",
				"start,HN,kj,HN,end",
				"start,HN,kj,dc,HN,end",
				"start,HN,kj,dc,end",
				"start,dc,HN,end",
				"start,dc,HN,kj,HN,end",
				"start,dc,end",
				"start,dc,kj,HN,end",
				"start,kj,HN,dc,HN,end",
				"start,kj,HN,dc,end",
				"start,kj,HN,end",
				"start,kj,dc,HN,end",
				"start,kj,dc,end",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []string
			g := readInput(tt.input)
			require.NotPanics(t, func() { result = path(g, "start", map[string]bool{}) })
			assert.ElementsMatch(t, tt.want, result)
		})
	}
}

func Test_path2(t *testing.T) {
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
	want := []string{
		"start,A,b,A,b,A,c,A,end",
		"start,A,b,A,b,A,end",
		"start,A,b,A,b,end",
		"start,A,b,A,c,A,b,A,end",
		"start,A,b,A,c,A,b,end",
		"start,A,b,A,c,A,c,A,end",
		"start,A,b,A,c,A,end",
		"start,A,b,A,end",
		"start,A,b,d,b,A,c,A,end",
		"start,A,b,d,b,A,end",
		"start,A,b,d,b,end",
		"start,A,b,end",
		"start,A,c,A,b,A,b,A,end",
		"start,A,c,A,b,A,b,end",
		"start,A,c,A,b,A,c,A,end",
		"start,A,c,A,b,A,end",
		"start,A,c,A,b,d,b,A,end",
		"start,A,c,A,b,d,b,end",
		"start,A,c,A,b,end",
		"start,A,c,A,c,A,b,A,end",
		"start,A,c,A,c,A,b,end",
		"start,A,c,A,c,A,end",
		"start,A,c,A,end",
		"start,A,end",
		"start,b,A,b,A,c,A,end",
		"start,b,A,b,A,end",
		"start,b,A,b,end",
		"start,b,A,c,A,b,A,end",
		"start,b,A,c,A,b,end",
		"start,b,A,c,A,c,A,end",
		"start,b,A,c,A,end",
		"start,b,A,end",
		"start,b,d,b,A,c,A,end",
		"start,b,d,b,A,end",
		"start,b,d,b,end",
		"start,b,end",
	}

	var result []string
	g := readInput(input)
	require.NotPanics(t, func() { result = path2(g, "start", map[string]bool{}, true) })
	r2 := []string{}
	found := map[string]bool{}
	for _, r := range result {
		if !found[r] {
			found[r] = true
			r2 = append(r2, r)
		}
	}
	assert.ElementsMatch(t, want, r2)
}
