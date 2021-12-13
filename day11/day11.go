package day11

import (
	"strings"

	"github.com/zsommers/aoc21/util"
)

type point struct{ x, y int }

func readInput(input []string) [][]int {
	m := [][]int{}
	for _, ss := range input {
		row := []int{}
		for _, s := range strings.Split(ss, "") {
			row = append(row, util.Atoi(s))
		}
		m = append(m, row)
	}
	return m
}

func flash(p point, m [][]int, flashed map[point]bool) {
	if flashed[p] {
		return
	}
	flashed[p] = true
	for x := util.Max(p.x-1, 0); x <= util.Min(p.x+1, len(m[0])-1); x++ {
		for y := util.Max(p.y-1, 0); y <= util.Min(p.y+1, len(m)-1); y++ {
			if x == p.x && y == p.y {
				continue
			}
			m[y][x]++
			if m[y][x] > 9 {
				flash(point{x, y}, m, flashed)
			}
		}
	}
}

func step(m [][]int) int {
	for x := 0; x < len(m[0]); x++ {
		for y := 0; y < len(m); y++ {
			m[y][x]++
		}
	}

	flashed := map[point]bool{}
	for x := 0; x < len(m[0]); x++ {
		for y := 0; y < len(m); y++ {
			if m[y][x] > 9 {
				flash(point{x, y}, m, flashed)
			}
		}
	}

	flashes := 0
	for x := 0; x < len(m[0]); x++ {
		for y := 0; y < len(m); y++ {
			if m[y][x] > 9 {
				m[y][x] = 0
				flashes++
			}
		}
	}

	return flashes
}

func steps(m [][]int, days int) int {
	sum := 0
	for d := 0; d < days; d++ {
		sum += step(m)
	}
	return sum
}

func A(input []string) int {
	return steps(readInput(input), 100)
}

func B(input []string) int {
	s := 0
	m := readInput(input)
	for {
		s++
		if step(m) == 100 {
			return s
		}
	}
}
