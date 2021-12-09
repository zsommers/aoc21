package day9

import (
	"github.com/zsommers/aoc21/util"
)

type point struct {
	x, y int
}

func readInput(input []string) [][]int {
	result := [][]int{}
	for _, rowI := range input {
		row := []int{}
		for _, c := range rowI {
			row = append(row, util.Atoi(string(c)))
		}
		result = append(result, row)
	}
	return result
}

func checkLowPoint(m [][]int, x, y int) bool {
	v := m[y][x]
	if x-1 >= 0 && m[y][x-1] <= v {
		return false
	}
	if x+1 < len(m[0]) && m[y][x+1] <= v {
		return false
	}
	if y-1 >= 0 && m[y-1][x] <= v {
		return false
	}
	if y+1 < len(m) && m[y+1][x] <= v {
		return false
	}
	return true
}

func findLowPoints(m [][]int) []point {
	result := []point{}
	for y := range m {
		for x := range m[y] {
			if checkLowPoint(m, x, y) {
				result = append(result, point{x, y})
			}
		}
	}
	return result
}

func A(input []string) int {
	m := readInput(input)
	sum := 0
	for _, p := range findLowPoints(m) {
		sum += 1 + m[p.y][p.x]
	}
	return sum
}

func notIn(p point, ps []point) bool {
	for _, pp := range ps {
		if p == pp {
			return false
		}
	}
	return true
}

func sizeBasin(m [][]int, x, y int, done []point) (int, []point) {
	done = append(done, point{x, y})
	sum := 1
	v := m[y][x]
	var n int
	if x-1 >= 0 {
		w := m[y][x-1]
		if w != 9 && w > v && notIn(point{x - 1, y}, done) {
			n, done = sizeBasin(m, x-1, y, done)
			sum += n
		}
	}
	if x+1 < len(m[0]) {
		w := m[y][x+1]
		if w != 9 && w > v && notIn(point{x + 1, y}, done) {
			n, done = sizeBasin(m, x+1, y, done)
			sum += n
		}
	}
	if y-1 >= 0 {
		w := m[y-1][x]
		if w != 9 && w > v && notIn(point{x, y - 1}, done) {
			n, done = sizeBasin(m, x, y-1, done)
			sum += n
		}
	}
	if y+1 < len(m) {
		w := m[y+1][x]
		if w != 9 && w > v && notIn(point{x, y + 1}, done) {
			n, done = sizeBasin(m, x, y+1, done)
			sum += n
		}
	}
	return sum, done
}

func B(input []string) int {
	m := readInput(input)
	var a, b, c int
	for _, p := range findLowPoints(m) {
		size, _ := sizeBasin(m, p.x, p.y, []point{})
		if size > a {
			c = b
			b = a
			a = size
		} else if size > b {
			c = b
			b = size
		} else if size > c {
			c = size
		}
	}

	return a * b * c
}
