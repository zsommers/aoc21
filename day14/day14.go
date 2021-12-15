package day14

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc21/util"
)

type pair struct{ a, b string }
type count map[string]int
type cacheKey struct {
	p pair
	c int
}
type cache map[cacheKey]count

func readInput(input []string) (string, map[pair]string) {
	insertions := map[pair]string{}
	for _, s := range input[2:] {
		ns := strings.Split(s, " -> ")
		if len(ns) != 2 {
			panic(fmt.Sprintf("cannot parse '%s'", s))
		}
		p := pair{string(ns[0][0]), string(ns[0][1])}
		insertions[p] = ns[1]
	}
	return input[0], insertions
}

func mergeCounts(cs ...count) count {
	result := map[string]int{}
	for _, c := range cs {
		for k, v := range c {
			if _, ok := result[k]; !ok {
				result[k] = v
			} else {
				result[k] += v
			}
		}
	}
	return result
}

func insert(p pair, rounds int, insertions map[pair]string, c cache) count {
	// Reached leaf
	if rounds == 1 {
		return count{
			insertions[p]: 1,
		}
	}

	// Check cache first
	if cnt, ok := c[cacheKey{p, rounds}]; ok {
		return cnt
	}

	// Recurse as last resort
	i := insertions[p]
	a := insert(pair{p.a, i}, rounds-1, insertions, c)
	b := insert(pair{i, p.b}, rounds-1, insertions, c)
	cnt := mergeCounts(a, b, count{i: 1})
	c[cacheKey{p, rounds}] = cnt
	return cnt
}

func polymerize(input []string, iterations int) int {
	polymer, insertions := readInput(input)
	sharedCache := cache{}
	mainCount := count{}

	p := pair{"", string(polymer[0])}
	templateCount := count{string(polymer[0]): 1}
	for _, c := range polymer[1:] {
		p.a = p.b
		p.b = string(c)
		if _, ok := templateCount[p.b]; !ok {
			templateCount[p.b] = 1
		} else {
			templateCount[p.b]++
		}

		cnt := insert(p, iterations, insertions, sharedCache)
		mainCount = mergeCounts(mainCount, cnt)
	}
	mainCount = mergeCounts(mainCount, templateCount)

	lowest := util.MaxInt
	highest := 0
	for _, v := range mainCount {
		lowest = util.Min(lowest, v)
		highest = util.Max(highest, v)
	}
	return highest - lowest
}

func A(input []string) int {
	return polymerize(input, 10)
}

func B(input []string) int {
	return polymerize(input, 40)
}
