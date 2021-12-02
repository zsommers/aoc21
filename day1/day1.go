package day1

import (
	"fmt"

	"github.com/zsommers/aoc21/util"
)

func A(input []string) int {
	increasing_count := 0

	// Get initial number
	last := util.Atoi(input[0])
	fmt.Printf("Intial value: %d\n", last)

	for _, s := range input[1:] {
		i := util.Atoi(s)
		increasing := i > last
		fmt.Printf("%d: Increasing? %t\n", i, increasing)
		if increasing {
			increasing_count += 1
		}
		last = i
	}

	return increasing_count
}

func B(input []string) int {
	increasing_count := 0

	// Get initial numbers
	first := util.Atoi(input[0])
	second := util.Atoi(input[1])
	third := util.Atoi(input[2])
	sum := first + second + third
	fmt.Printf("Intial values: %d %d %d - %d\n", first, second, third, sum)

	for _, s := range input[3:] {
		first = second
		second = third
		third = util.Atoi(s)
		inner_sum := first + second + third
		increasing := inner_sum > sum
		fmt.Printf("%d %d %d - %d: Increasing? %t\n", first, second, third, inner_sum, increasing)
		if increasing {
			increasing_count += 1
		}
		sum = inner_sum
	}

	return increasing_count
}
