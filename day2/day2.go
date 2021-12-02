package day2

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc21/util"
)

func A(input []string) int {
	distance := 0
	depth := 0

	for _, s := range input {
		parts := strings.Split(s, " ")
		i := util.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			distance += i
		case "down":
			depth += i
		case "up":
			depth -= i
		default:
			panic(fmt.Sprintf("Unknown command '%s'", parts[0]))
		}
	}
	return distance * depth
}

func B(input []string) int {
	distance := 0
	depth := 0
	aim := 0

	for _, s := range input {
		parts := strings.Split(s, " ")
		i := util.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			distance += i
			depth += aim * i
		case "down":
			aim += i
		case "up":
			aim -= i
		default:
			panic(fmt.Sprintf("Unknown command '%s'", parts[0]))
		}
	}
	return distance * depth
}
