package day10

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 26397, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 288957, B(input))
}

func Test_stack_Push(t *testing.T) {
	tests := []struct {
		s    stack
		e    string
		want stack
	}{
		{
			s:    stack{},
			e:    "(",
			want: stack{"("},
		},
		{
			s:    stack{"["},
			e:    "(",
			want: stack{"[", "("},
		},
	}
	for _, tt := range tests {
		t.Run(tt.e, func(t *testing.T) {
			assert.NotPanics(t, func() { tt.s.Push(tt.e) })
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	tests := []struct {
		s      stack
		wantE  string
		wantOk bool
		wantS  stack
	}{
		{
			s:      stack{},
			wantE:  "",
			wantOk: false,
			wantS:  stack{},
		},
		{
			s:      stack{"["},
			wantE:  "[",
			wantOk: true,
			wantS:  stack{},
		},
		{
			s:      stack{"[", "("},
			wantE:  "(",
			wantOk: true,
			wantS:  stack{"["},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.s), func(t *testing.T) {
			var e string
			var ok bool
			assert.NotPanics(t, func() { e, ok = tt.s.Pop() })
			assert.Equal(t, tt.wantE, e)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantS, tt.s)
		})
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		input string
		wantS string
		wantE error
	}{
		{
			input: "",
			wantE: nil,
		},
		{
			input: "[({(<(())[]>[[{[]{<()<>>",
			wantS: "}}]])})]",
			wantE: nil,
		},
		{
			input: "[(()[<>])]({[<{<<[]>>(",
			wantS: ")}>]})",
			wantE: nil,
		},
		{
			input: "(((({<>}<{<{<>}{[]{[]{}",
			wantS: "}}>}>))))",
			wantE: nil,
		},
		{
			input: "{<[[]]>}<{[{[{[]{()[[[]",
			wantS: "]]}}]}]}>",
			wantE: nil,
		},
		{
			input: "<{([{{}}[<[[[<>{}]]]>[]]",
			wantS: "])}>",
			wantE: nil,
		},
		{
			input: "{([(<{}[<>[]}>{[]{[(<()>",
			wantS: "",
			wantE: parseError{12, "}", "]"},
		},
		{
			input: "[[<[([]))<([[{}[[()]]]",
			wantS: "",
			wantE: parseError{8, ")", "]"},
		},
		{
			input: "[{[{({}]{}}([{[{{{}}([]",
			wantS: "",
			wantE: parseError{7, "]", ")"},
		},
		{
			input: "[<(<(<(<{}))><([]([]()",
			wantS: "",
			wantE: parseError{10, ")", ">"},
		},
		{
			input: "<{([([[(<>()){}]>(<<{{",
			wantS: "",
			wantE: parseError{16, ">", "]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			fix, err := parse(tt.input)
			assert.Equal(t, tt.wantS, fix)
			assert.ErrorIs(t, err, tt.wantE)
		})
	}
}
