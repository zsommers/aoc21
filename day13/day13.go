package day13

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc21/util"
)

type point struct{ x, y int }
type fold struct {
	axis string
	idx  int
}

func readInput(input []string) (map[point]bool, []fold) {
	points := map[point]bool{}
	var idx int
	var s string
	for idx, s = range input {
		if s == "" {
			break
		}
		ns := strings.Split(s, ",")
		if len(ns) != 2 {
			panic(fmt.Sprintf("cannot parse '%s'", s))
		}
		p := point{util.Atoi(ns[0]), util.Atoi(ns[1])}
		points[p] = true
	}
	folds := []fold{}
	for _, s := range input[idx+1:] {
		ns := strings.Split(s[11:], "=")
		if len(ns) != 2 {
			panic(fmt.Sprintf("cannot parse '%s'", s))
		}
		folds = append(folds, fold{ns[0], util.Atoi(ns[1])})
	}
	return points, folds
}

func foldX(points map[point]bool, f fold) map[point]bool {
	newPoints := map[point]bool{}
	for p := range points {
		if p.x < f.idx {
			newPoints[p] = true
		} else {
			np := point{2*f.idx - p.x, p.y}
			newPoints[np] = true
		}
	}
	return newPoints
}

func foldY(points map[point]bool, f fold) map[point]bool {
	newPoints := map[point]bool{}
	for p := range points {
		if p.y < f.idx {
			newPoints[p] = true
		} else {
			np := point{p.x, 2*f.idx - p.y}
			newPoints[np] = true
		}
	}
	return newPoints
}

func foldOne(points map[point]bool, f fold) map[point]bool {
	if f.axis == "x" {
		return foldX(points, f)
	} else {
		return foldY(points, f)
	}
}

func print(points map[point]bool) {
	var maxX, maxY int
	for p := range points {
		maxX = util.Max(maxX, p.x)
		maxY = util.Max(maxY, p.y)
	}
	fmt.Println(points)
	grid := [][]string{}
	for i := 0; i <= maxY; i++ {
		row := []string{}
		for j := 0; j <= maxX; j++ {
			row = append(row, " ")
		}
		grid = append(grid, row)
	}
	for p := range points {
		grid[p.y][p.x] = "#"
	}
	fmt.Println()
	for _, r := range grid {
		fmt.Println(strings.Join(r, ""))
	}
	fmt.Println()
}

func A(input []string) int {
	points, folds := readInput(input)
	points = foldOne(points, folds[0])
	return len(points)
}

func B(input []string) int {
	points, folds := readInput(input)
	for _, f := range folds {
		points = foldOne(points, f)
	}
	print(points)
	return len(points)
}
