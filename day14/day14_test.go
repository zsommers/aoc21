package day14

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

var rawInput2 = `PPFCHPFNCKOKOSBVCFPP

VC -> N
SC -> H
CK -> P
OK -> O
KV -> O
HS -> B
OH -> O
VN -> F
FS -> S
ON -> B
OS -> H
PC -> B
BP -> O
OO -> N
BF -> K
CN -> B
FK -> F
NP -> K
KK -> H
CB -> S
CV -> K
VS -> F
SF -> N
KB -> H
KN -> F
CP -> V
BO -> N
SS -> O
HF -> H
NN -> F
PP -> O
VP -> H
BB -> K
VB -> N
OF -> N
SH -> S
PO -> F
OC -> S
NS -> C
FH -> N
FP -> C
SO -> P
VK -> C
HP -> O
PV -> S
HN -> K
NB -> C
NV -> K
NK -> B
FN -> C
VV -> N
BN -> N
BH -> S
FO -> V
PK -> N
PS -> O
CO -> K
NO -> K
SV -> C
KO -> V
HC -> B
BC -> N
PB -> C
SK -> S
FV -> K
HO -> O
CF -> O
HB -> P
SP -> N
VH -> P
NC -> K
KC -> B
OV -> P
BK -> F
FB -> F
FF -> V
CS -> F
CC -> H
SB -> C
VO -> V
VF -> O
KP -> N
HV -> H
PF -> H
KH -> P
KS -> S
BS -> H
PH -> S
SN -> K
HK -> P
FC -> N
PN -> S
HH -> N
OB -> P
BV -> S
KF -> N
OP -> H
NF -> V
CH -> K
NH -> P`

var input = strings.Split(rawInput, "\n")
var input2 = strings.Split(rawInput2, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 1588, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 2188189693529, B(input))
}

func Test_readInput(t *testing.T) {
	wantS := "NNCB"
	wantM := map[pair]string{
		{"C", "H"}: "B",
		{"H", "H"}: "N",
		{"C", "B"}: "H",
		{"N", "H"}: "C",
		{"H", "B"}: "C",
		{"H", "C"}: "B",
		{"H", "N"}: "C",
		{"N", "N"}: "C",
		{"B", "H"}: "H",
		{"N", "C"}: "B",
		{"N", "B"}: "B",
		{"B", "N"}: "B",
		{"B", "B"}: "N",
		{"B", "C"}: "B",
		{"C", "C"}: "N",
		{"C", "N"}: "C",
	}
	var s string
	var m map[pair]string
	require.NotPanics(t, func() { s, m = readInput(input) })
	assert.Equal(t, wantS, s)
	assert.Equal(t, wantM, m)
}

func Test_insert(t *testing.T) {
	tests := []struct {
		rounds int
		want   count
	}{
		{
			1,
			// "NCNBCHB",
			count{
				"N": 2,
				"C": 2,
				"B": 2,
				"H": 1,
			},
		},
		{
			2,
			// "NBCCNBBBCBHCB",
			count{
				"N": 2,
				"C": 4,
				"B": 6,
				"H": 1,
			},
		},
		{
			3,
			// "NBBBCNCCNBBNBNBBCHBHHBCHB",
			count{
				"N": 5,
				"C": 5,
				"B": 11,
				"H": 4,
			},
		},
		{
			4,
			// "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB",
			count{
				"N": 11,
				"C": 10,
				"B": 23,
				"H": 5,
			},
		},
	}
	_, insertions := readInput(input)
	templateCount := count{
		"N": 2,
		"C": 1,
		"B": 1,
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("round %d", tt.rounds), func(t *testing.T) {
			a := insert(pair{"N", "N"}, tt.rounds, insertions, cache{})
			b := insert(pair{"N", "C"}, tt.rounds, insertions, cache{})
			c := insert(pair{"C", "B"}, tt.rounds, insertions, cache{})
			assert.Equal(t, tt.want, mergeCounts(templateCount, a, b, c))
		})
	}
}

func Test_cache(t *testing.T) {
	tests := []struct {
		pair   pair
		rounds int
		cache  cache
		want   count
	}{
		{
			pair{"A", "B"},
			4,
			cache{
				cacheKey{pair{"A", "B"}, 4}: count{"C": 7},
			},
			count{
				"A": 1,
				"B": 1,
				"C": 7,
			},
		},
		{
			pair{"N", "C"},
			5,
			cache{
				cacheKey{pair{"N", "B"}, 4}: count{"C": 7},
				cacheKey{pair{"B", "C"}, 4}: count{"N": 7},
			},
			count{
				"N": 8,
				"B": 1,
				"C": 8,
			},
		},
		{
			pair{"N", "C"},
			5,
			cache{
				cacheKey{pair{"N", "B"}, 4}: count{"C": 7},
				cacheKey{pair{"B", "B"}, 3}: count{"N": 3},
			},
			count{
				"N": 6,
				"B": 7,
				"C": 8,
			},
		},
	}
	_, insertions := readInput(input)
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s - %d", tt.pair.a, tt.pair.b, tt.rounds), func(t *testing.T) {
			cnt := count{tt.pair.a: 1, tt.pair.b: 1}
			result := insert(tt.pair, tt.rounds, insertions, tt.cache)
			assert.Equal(t, tt.want, mergeCounts(cnt, result))
		})
	}
}

func Test_insert2(t *testing.T) {
	tests := []struct {
		rounds int
		want   count
	}{
		{
			1,
			// "POP",
			count{
				"P": 2,
				"O": 1,
			},
		},
		{
			2,
			// "PFOHP",
			count{
				"P": 2,
				"O": 1,
				"F": 1,
				"H": 1,
			},
		},
		{
			3,
			// "PHFVOOHOP",
			count{
				"P": 2,
				"O": 3,
				"F": 1,
				"H": 2,
				"V": 1,
			},
		},
		{
			4,
			// "PSHHFKVVONOOHOOHP",
			count{
				"P": 2,
				"O": 5,
				"F": 1,
				"H": 4,
				"V": 2,
				"S": 1,
				"K": 1,
				"N": 1,
			},
		},
	}
	_, insertions := readInput(input2)
	templateCount := count{
		"P": 2,
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("round %d", tt.rounds), func(t *testing.T) {
			a := insert(pair{"P", "P"}, tt.rounds, insertions, cache{})
			assert.Equal(t, tt.want, mergeCounts(templateCount, a))
		})
	}
}
