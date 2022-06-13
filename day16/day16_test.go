package day16

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = ``

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	tests := []struct {
		hex  string
		want int
	}{
		{
			"8A004A801A8002F478",
			16,
		},
		{
			"620080001611562C8802118E34",
			12,
		},
		{
			"C0015000016115A2E0802F182340",
			23,
		},
		{
			"A0016C880162017C3686B18A3D4780",
			31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			var got int
			require.NotPanics(t, func() { got = A([]string{tt.hex}) })
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestB(t *testing.T) {
	tests := []struct {
		hex  string
		want int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			var got int
			require.NotPanics(t, func() { got = B([]string{tt.hex}) })
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_readInput(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"D2FE28",
			"110100101111111000101000",
		},
		{
			"38006F45291200",
			"00111000000000000110111101000101001010010001001000000000",
		},
		{
			"EE00D40C823060",
			"11101110000000001101010000001100100000100011000001100000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			var bits string
			require.NotPanics(t, func() { bits = readInput(tt.s) })
			assert.Equal(t, tt.want, bits)
		})
	}
}

func Test_parsePacket(t *testing.T) {
	tests := []struct {
		hex   string
		wantP packet
		wantB int
	}{
		{
			"D2FE28",
			packet{
				version: 6,
				typeID:  4,
				value:   2021,
			},
			21,
		},
		{
			"38006F45291200",
			packet{
				version: 1,
				typeID:  6,
				subPackets: []packet{
					{
						version: 6,
						typeID:  4,
						value:   10,
					},
					{
						version: 2,
						typeID:  4,
						value:   20,
					},
				},
			},
			49,
		},
		{
			"EE00D40C823060",
			packet{
				version: 7,
				typeID:  3,
				subPackets: []packet{
					{
						version: 2,
						typeID:  4,
						value:   1,
					},
					{
						version: 4,
						typeID:  4,
						value:   2,
					},
					{
						version: 1,
						typeID:  4,
						value:   3,
					},
				},
			},
			51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			var p packet
			var b int
			bits := readInput(tt.hex)
			require.NotPanics(t, func() { p, b = parsePacket(bits) })
			assert.Equal(t, tt.wantP, p)
			assert.Equal(t, tt.wantB, b)
		})
	}
}
