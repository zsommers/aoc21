package day16

import (
	"fmt"
	"strconv"

	"github.com/zsommers/aoc21/util"
)

type packet struct {
	version    uint8
	typeID     uint8
	value      uint64
	subPackets []packet
}

func readInput(s string) string {
	bits := ""
	for _, c := range s {
		bit, err := strconv.ParseUint(string(c), 16, 8)
		util.CheckErr(err)
		bits += fmt.Sprintf("%04b", bit)
	}
	return bits
}

func parseValue(bits string) (uint64, int) {
	var val uint64
	var idx int
	for ; idx < len(bits); idx += 5 {
		val = val << 4
		v, err := strconv.ParseUint(bits[idx+1:idx+5], 2, 64)
		util.CheckErr(err)
		val += v
		if bits[idx] == '0' {
			break
		}
	}
	return val, idx + 5
}

func parseSubPackets(bits string) ([]packet, int) {
	ps := []packet{}

	if bits[0] == '0' {
		length, err := strconv.ParseUint(bits[1:16], 2, 0)
		util.CheckErr(err)
		bitsRead := 16
		for uint64(bitsRead-16) < length {
			p, b := parsePacket(bits[bitsRead : 16+length])
			ps = append(ps, p)
			bitsRead += b
		}
		return ps, bitsRead
	}

	pCnt, err := strconv.ParseUint(bits[1:12], 2, 0)
	util.CheckErr(err)
	bitsRead := 12
	for i := 0; i < int(pCnt); i++ {
		p, b := parsePacket(bits[bitsRead:])
		ps = append(ps, p)
		bitsRead += b
	}

	return ps, bitsRead
}

func parsePacket(bits string) (packet, int) {
	p := packet{}

	v, err := strconv.ParseUint(bits[0:3], 2, 8)
	util.CheckErr(err)
	p.version = uint8(v)

	t, err := strconv.ParseUint(bits[3:6], 2, 8)
	util.CheckErr(err)
	p.typeID = uint8(t)

	bitsRead := 6

	b := 0
	if p.typeID != 4 {
		p.subPackets, b = parseSubPackets(bits[6:])
		bitsRead += b
	} else {
		p.value, b = parseValue(bits[6:])
		bitsRead += b
	}

	return p, bitsRead
}

func sumVersions(p packet) int {
	sum := int(p.version)
	for _, pp := range p.subPackets {
		sum += sumVersions(pp)
	}
	return sum
}

func A(input []string) int {
	bits := readInput(input[0])
	p, _ := parsePacket(bits)
	return sumVersions(p)
}

func evaluate(p packet) int {
	switch p.typeID {
	case 0:
		sum := 0
		for _, pp := range p.subPackets {
			sum += evaluate(pp)
		}
		return sum
	case 1:
		product := 1
		for _, pp := range p.subPackets {
			product *= evaluate(pp)
		}
		return product
	case 2:
		values := []int{}
		for _, pp := range p.subPackets {
			values = append(values, evaluate(pp))
		}
		return util.Min(values...)
	case 3:
		values := []int{}
		for _, pp := range p.subPackets {
			values = append(values, evaluate(pp))
		}
		return util.Max(values...)
	case 4:
		return int(p.value)
	case 5:
		if evaluate(p.subPackets[0]) > evaluate(p.subPackets[1]) {
			return 1
		}
		return 0
	case 6:
		if evaluate(p.subPackets[0]) < evaluate(p.subPackets[1]) {
			return 1
		}
		return 0
	case 7:
		if evaluate(p.subPackets[0]) == evaluate(p.subPackets[1]) {
			return 1
		}
		return 0
	default:
		panic(fmt.Sprintf("cannot evaluate %v", p))
	}
}

func B(input []string) int {
	bits := readInput(input[0])
	p, _ := parsePacket(bits)
	return evaluate(p)
}
