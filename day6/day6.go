package day6

import (
	"strings"

	"github.com/zsommers/aoc21/util"
)

func update(old []int) []int {
	new := make([]int, 9)
	for i := 7; i >= 0; i-- {
		new[i] = old[i+1]
	}
	new[6] += old[0]
	new[8] = old[0]

	return new
}

/*func update(state, days int) int {
	if state >= days {
		return 1
	}

	return update(6, days-state-1) + update(8, days-state-1)
}*/

func readInput(s string) []int {
	result := make([]int, 9)
	for _, raw := range strings.Split(s, ",") {
		result[util.Atoi(raw)]++
	}
	return result
}

func A(input []string) int {
	state := readInput(input[0])
	for i := 0; i < 80; i++ {
		state = update(state)
	}

	sum := 0
	for _, c := range state {
		sum += c
	}

	return sum
}

func B(input []string) int {
	state := readInput(input[0])
	for i := 0; i < 256; i++ {
		state = update(state)
	}

	sum := 0
	for _, c := range state {
		sum += c
	}

	return sum
}
