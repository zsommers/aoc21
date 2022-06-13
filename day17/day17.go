package day17

import (
	"regexp"

	"github.com/zsommers/aoc21/util"
)

type bounds struct{ minX, maxX, minY, maxY int }

func (b *bounds) contains(posX, posY int) bool {
	return (posX >= b.minX &&
		posX <= b.maxX &&
		posY >= b.minY &&
		posY <= b.maxY)
}

func readInput(s string) bounds {
	reX := regexp.MustCompile(`x=(-?\d+)..(-?\d+)`)
	reY := regexp.MustCompile(`y=(-?\d+)..(-?\d+)`)
	x1 := util.Atoi(reX.FindStringSubmatch(s)[1])
	x2 := util.Atoi(reX.FindStringSubmatch(s)[2])
	y1 := util.Atoi(reY.FindStringSubmatch(s)[1])
	y2 := util.Atoi(reY.FindStringSubmatch(s)[2])

	b := bounds{}
	b.minX = util.Min(x1, x2)
	b.maxX = util.Max(x1, x2)
	b.minY = util.Min(y1, y2)
	b.maxY = util.Max(y1, y2)

	return b
}

func x(initialV, time int) int {
	v := time * initialV
	drag := 0
	if time <= initialV {
		drag = time * (time - 1) / 2
	} else {
		drag = initialV * (2*time - 7) / 2
	}
	return v - drag
}

func y(initialV, time int) int {
	v := time * initialV
	g := time * (time - 1) / 2
	return v - g
}

func A(input []string) int {
	b := readInput(input[0])
	initialV := -(b.minY + 1)
	return y(initialV, initialV+1)
}

func checkSuccess(velX, velY int, b bounds) bool {
	n := 1
	var posX, posY int
	for posX < b.maxX && posY > b.minY {
		posX = x(velX, n)
		posY = y(velY, n)

		if b.contains(posX, posY) {
			return true
		}
		n++
	}
	return false
}

func B(input []string) int {
	cnt := 0
	b := readInput(input[0])
	for xV := 1; xV <= b.maxX; xV++ {
		for yV := b.minY; yV <= -(b.minY + 1); yV++ {
			if checkSuccess(xV, yV, b) {
				// fmt.Printf("%d %d\n", xV, yV)
				cnt++
			}
		}
	}
	return cnt
}
