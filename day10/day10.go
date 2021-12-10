package day10

import (
	"fmt"
	"sort"
)

type parseError struct {
	idx      int
	found    string
	expected string
}

func (e parseError) Error() string {
	return fmt.Sprintf("Parse error at index %d: expected %s, found %s", e.idx, e.expected, e.found)
}

type stack []string

func (s *stack) Push(e string) {
	*s = append(*s, e)
}

func (s *stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}

	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e, true
}

func close(open string) string {
	switch open {
	case "(":
		return ")"
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	default:
		panic(fmt.Sprintf("Unexpected open %s", open))
	}
}

func parse(input string) (string, error) {
	s := stack{}
	for i, r := range input {
		switch el := string(r); el {
		case "(":
			fallthrough
		case "{":
			fallthrough
		case "[":
			fallthrough
		case "<":
			s.Push(el)
		case ")":
			prior, ok := s.Pop()
			if !ok || prior != "(" {
				return "", parseError{i, el, close(prior)}
			}
		case "}":
			if prior, ok := s.Pop(); !ok || prior != "{" {
				return "", parseError{i, el, close(prior)}
			}
		case "]":
			if prior, ok := s.Pop(); !ok || prior != "[" {
				return "", parseError{i, el, close(prior)}
			}
		case ">":
			if prior, ok := s.Pop(); !ok || prior != "<" {
				return "", parseError{i, el, close(prior)}
			}
		default:
			panic(fmt.Sprintf("invalid character %s at %d", el, i))
		}
	}

	fix := ""
	for el, ok := s.Pop(); ok; {
		fix += close(el)
		el, ok = s.Pop()
	}
	return fix, nil
}

func A(input []string) int {
	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	score := 0
	for _, s := range input {
		if _, err := parse(s); err != nil {
			c := (err.(parseError)).found
			score += points[c]
		}
	}
	return score
}

func B(input []string) int {
	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	scores := []int{}
	for _, s := range input {
		score := 0
		fix, err := parse(s)
		if err != nil || fix == "" {
			continue
		}
		for _, c := range fix {
			score *= 5
			score += points[string(c)]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}
