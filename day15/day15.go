package day15

import (
	"fmt"

	"github.com/zsommers/aoc21/util"
)

type point struct{ x, y int }

func readInput(input []string) [][]int {
	g := [][]int{}
	for _, s := range input {
		r := []int{}
		for _, c := range s {
			r = append(r, util.Atoi(string(c)))
		}
		g = append(g, r)
	}
	return g
}

func lowest(m map[point]int) (point, int) {
	lr := util.MaxInt
	var lp point
	for p, r := range m {
		if r < lr {
			lr = r
			lp = p
		}
	}
	return lp, lr
}

func visit(p point, g [][]int, unvisted map[point]int) {
	neighbors := []point{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
	}
	for _, n := range neighbors {
		// Skip visited nodes
		if _, ok := unvisted[n]; !ok {
			continue
		}

		// Update neighbor cost if lower
		newR := unvisted[p] + g[n.y][n.x]
		if newR < unvisted[n] {
			unvisted[n] = newR
		}
	}

	delete(unvisted, p)
	fmt.Printf("Removed { %3d %3d } - %6d unvisited\n", p.x, p.y, len(unvisted))
}

func makeUnvisited(width, height int) map[point]int {
	g := map[point]int{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			g[point{x, y}] = util.MaxInt - 1
		}
	}
	return g
}

func traverse(g [][]int, unvisited map[point]int) {
	p := point{0, 0}
	unvisited[p] = 0
	for {
		visit(p, g, unvisited)
		p, _ = lowest(unvisited)
		if p.x == len(g[0])-1 && p.y == len(g)-1 {
			// path found
			return
		}
	}
}

func embiggen(g [][]int) [][]int {
	big := make([][]int, len(g)*5)
	for i, r := range g {
		big[i] = append(big[i], r...)
	}
	// Build top row
	for cnt := 1; cnt < 5; cnt++ {
		for i, r := range g {
			for _, n := range r {
				n += cnt
				if n > 9 {
					n -= 9
				}
				big[i] = append(big[i], n)
			}
		}
	}
	// Build next 4
	for cntY := 1; cntY < 5; cntY++ {
		// Copy prior row
		offset := len(g) * cntY
		for i, r := range big[offset-len(g) : offset] {
			big[offset+i] = append(big[offset+i], r[len(g[0]):]...)
		}
		// Extend new row
		for i, r := range big[offset : offset+len(g)] {
			for _, n := range r[len(r)-len(g[0]):] {
				n++
				if n > 9 {
					n -= 9
				}
				big[offset+i] = append(big[offset+i], n)
			}
		}
	}
	return big
}

func A(input []string) int {
	g := readInput(input)
	unvisited := makeUnvisited(len(g[0]), len(g))
	traverse(g, unvisited)
	_, r := lowest(unvisited)
	return r
}

func B(input []string) int {
	g := readInput(input)
	g = embiggen(g)
	unvisited := makeUnvisited(len(g[0]), len(g))
	traverse(g, unvisited)
	_, r := lowest(unvisited)
	return r
}
