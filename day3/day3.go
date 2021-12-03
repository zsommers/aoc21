package day3

import (
	"fmt"
	"strconv"

	"github.com/zsommers/aoc21/util"
)

func countBits(input []string) []map[byte]int {
	counts := []map[byte]int{}
	for i := 0; i < len(input[0]); i += 1 {
		counts = append(counts, map[byte]int{
			'0': 0,
			'1': 0,
		})
	}
	for _, s := range input {
		for i, c := range s {
			counts[i][byte(c)] += 1
		}
	}
	return counts
}

func A(input []string) int {
	counts := countBits(input)

	var gamma, epsilon string
	for _, c := range counts {
		if c['1'] > c['0'] {
			gamma += "1"
			epsilon += "0"
		} else if c['1'] < c['0'] {
			gamma += "0"
			epsilon += "1"
		} else {
			panic("Equal counts!")
		}
	}

	gammaI, err := strconv.ParseInt(gamma, 2, 64)
	util.CheckErr(err)
	epsilonI, err := strconv.ParseInt(epsilon, 2, 64)
	util.CheckErr(err)
	return int(gammaI * epsilonI)
}

func filterList(input []string, index int, match byte) []string {
	output := []string{}
	for _, i := range input {
		if i[index] == match {
			output = append(output, i)
		}
	}

	return output
}

func B(input []string) int {

	o2 := make([]string, len(input))
	copy(o2, input)
	for i := range o2[0] {
		counts := countBits(o2)
		match := '1'
		if counts[i]['0'] > counts[i]['1'] {
			match = '0'
		}
		o2 = filterList(o2, i, byte(match))
		if len(o2) == 1 {
			break
		}
	}

	co2 := make([]string, len(input))
	copy(co2, input)
	for i := range co2[0] {
		counts := countBits(co2)
		match := '0'
		if counts[i]['1'] < counts[i]['0'] {
			match = '1'
		}
		co2 = filterList(co2, i, byte(match))
		if len(co2) == 1 {
			break
		}
	}

	if len(o2) != 1 || len(co2) != 1 {
		fmt.Println(o2)
		fmt.Println(co2)
		panic("")
	}

	o2Int, err := strconv.ParseInt(o2[0], 2, 64)
	util.CheckErr(err)
	co2Int, err := strconv.ParseInt(co2[0], 2, 64)
	util.CheckErr(err)
	return int(o2Int * co2Int)
}
