package day7

import (
	"github.com/zsommers/aoc21/util"
)

func costA(positions []int, destination int) int {
	cost := 0
	for _, p := range positions {
		cost += util.Abs(destination - p)
	}
	return cost
}

func costB(positions []int, destination int) int {
	cost := 0
	for _, p := range positions {
		d := util.Abs(destination - p)
		cost += d * (d + 1) / 2
	}
	return cost
}

func A(input []string) int {
	positions := util.ReadIntString(input[0])
	var max, min int
	for _, p := range positions {
		max = util.Max(max, p)
		min = util.Min(min, p)
	}
	c := int(^uint(0) >> 1)
	for i := min; i < max; i++ {
		c = util.Min(c, costA(positions, i))
	}
	return c
}

func B(input []string) int {
	positions := util.ReadIntString(input[0])
	var max, min int
	for _, p := range positions {
		max = util.Max(max, p)
		min = util.Min(min, p)
	}
	c := int(^uint(0) >> 1)
	for i := min; i < max; i++ {
		c = util.Min(c, costB(positions, i))
	}
	return c
}
