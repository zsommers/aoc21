package day8

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func split(s string) ([]string, []string) {
	ss := strings.Split(s, " | ")
	if len(ss) != 2 {
		panic(fmt.Sprintf("invalid input: %s", s))
	}
	return strings.Fields(ss[0]), strings.Fields(ss[1])
}

func A(input []string) int {
	count := 0
	for _, s := range input {
		_, output := split(s)
		for _, o := range output {
			switch len(o) {
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 7:
				count++
			}
		}
	}
	return count
}

func fourLeft(one, four string) string {
	left := ""
	for _, c := range four {
		if !strings.Contains(one, string(c)) {
			left += string(c)
		}
	}
	return left
}

func containsSegments(digit, segments string, match int) bool {
	count := 0
	for _, c := range segments {
		if strings.Contains(digit, string(c)) {
			count++
		}
	}
	return count == match
}

func lenFives(lenFive map[string]bool, digits []string) {
	for d := range lenFive {
		if containsSegments(d, fourLeft(digits[1], digits[4]), 2) {
			digits[5] = d
		} else if containsSegments(d, digits[1], 2) {
			digits[3] = d
		} else {
			digits[2] = d
		}
	}
}

func lenSixes(lenSix map[string]bool, digits []string) {
	for d := range lenSix {
		if containsSegments(d, fourLeft(digits[1], digits[4]), 1) {
			digits[0] = d
		} else if containsSegments(d, digits[1], 1) {
			digits[6] = d
		} else {
			digits[9] = d
		}
	}
}

func mapDigits(ss []string) map[string]int {
	digits := make([]string, 10)
	lenFive := map[string]bool{}
	lenSix := map[string]bool{}
	for _, digit := range ss {
		switch len(digit) {
		case 2:
			digits[1] = digit
		case 3:
			digits[7] = digit
		case 4:
			digits[4] = digit
		case 5:
			lenFive[digit] = true
		case 6:
			lenSix[digit] = true
		case 7:
			digits[8] = digit
		}
	}
	if digits[1] == "" || digits[4] == "" || digits[7] == "" || digits[8] == "" {
		panic(fmt.Sprintf("missing digit %s", ss))
	}

	lenFives(lenFive, digits)
	lenSixes(lenSix, digits)

	digitMap := map[string]int{}
	for i, d := range digits {
		digitMap[sortString(d)] = i
	}
	return digitMap
}

func sortString(s string) string {
	cs := strings.Split(s, "")
	sort.Strings(cs)
	return strings.Join(cs, "")
}

func sumDigits(output []string, digitMap map[string]int) int {
	result := 0
	for i, o := range output {
		n := digitMap[sortString(o)]
		ex := int(math.Pow(10, float64(len(output)-i-1)))
		result += n * ex
	}
	return result
}

func B(input []string) int {
	result := 0
	for _, s := range input {
		input, output := split(s)
		digitMap := mapDigits(append(input, output...))
		result += sumDigits(output, digitMap)
	}
	return result
}
