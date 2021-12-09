package day8

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 26, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 61229, B(input))
}

func Test_containsSegments(t *testing.T) {
	tests := []struct {
		digit    string
		segments string
		match    int
		want     bool
	}{
		{
			digit:    "abc",
			segments: "bc",
			match:    2,
			want:     true,
		},
		{
			digit:    "abc",
			segments: "bc",
			match:    1,
			want:     false,
		},
		{
			digit:    "abc",
			segments: "cd",
			match:    2,
			want:     false,
		},
		{
			digit:    "abc",
			segments: "cd",
			match:    1,
			want:     true,
		},
		{
			digit:    "abc",
			segments: "",
			match:    0,
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s %d", tt.digit, tt.segments, tt.match), func(t *testing.T) {
			assert.Equal(t, tt.want, containsSegments(tt.digit, tt.segments, tt.match))
		})
	}
}

func Test_fourLeft(t *testing.T) {
	assert.Equal(t, "cd", fourLeft("ab", "abcd"))
}

func Test_mapDigits(t *testing.T) {
	input := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	expected := map[string]int{
		"abcdefg": 8,
		"bcdef":   5,
		"acdfg":   2,
		"abcdf":   3,
		"abd":     7,
		"abcdef":  9,
		"bcdefg":  6,
		"abef":    4,
		"abcdeg":  0,
		"ab":      1,
	}
	assert.Equal(t, expected, mapDigits(strings.Fields(input)))
}

func Test_sumDigits(t *testing.T) {
	rawInput := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	input, output := split(rawInput)
	dm := mapDigits(append(input, output...))
	assert.Equal(t, 5353, sumDigits(output, dm))
}
