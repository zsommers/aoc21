package day5

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc21/util"
)

type point struct {
	x int
	y int
}

func readPoint(s string) point {
	ns := strings.Split(s, ",")
	if len(ns) != 2 {
		panic(fmt.Sprintf("%s resulted in %d coords", s, len(ns)))
	}
	return point{util.Atoi(ns[0]), util.Atoi(ns[1])}
}

func readSegment(s string) (a, b point) {
	points := strings.Split(s, " -> ")
	if len(points) != 2 {
		panic(fmt.Sprintf("%s resulted in %d points", s, len(points)))
	}
	return readPoint(points[0]), readPoint(points[1])
}

func createRange(a, b point) []point {
	result := []point{}
	var xMod, yMod int
	if a.x < b.x {
		xMod = 1
	} else if a.x > b.x {
		xMod = -1
	}
	if a.y < b.y {
		yMod = 1
	} else if a.y > b.y {
		yMod = -1
	}

	l := util.Max(util.Abs(a.x-b.x), util.Abs(a.y-b.y))
	x := a.x
	y := a.y

	for i := 0; i <= l; i++ {
		result = append(result, point{x, y})
		x += xMod
		y += yMod
	}

	return result
}

func A(input []string) int {
	vents := map[point]int{}

	for _, s := range input {
		a, b := readSegment(s)
		if !(a.x == b.x || a.y == b.y) {
			continue
		}
		for _, p := range createRange(a, b) {
			if _, ok := vents[p]; !ok {
				vents[p] = 1
			} else {
				vents[p]++
			}
		}
	}

	total := 0
	for _, n := range vents {
		if n >= 2 {
			total++
		}
	}

	return total
}

func B(input []string) int {
	vents := map[point]int{}

	for _, s := range input {
		a, b := readSegment(s)
		for _, p := range createRange(a, b) {
			if _, ok := vents[p]; !ok {
				vents[p] = 1
			} else {
				vents[p]++
			}
		}
	}

	total := 0
	for _, n := range vents {
		if n >= 2 {
			total++
		}
	}

	return total
}
