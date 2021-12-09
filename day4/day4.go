package day4

import (
	"strings"

	"github.com/zsommers/aoc21/util"
)

type boardSpace struct {
	number int
	called bool
}

type board [][]boardSpace

func (b board) mark(draw int) bool {
	for r := range b {
		for c := range b[r] {
			if b[r][c].number == draw {
				b[r][c].called = true
				return true
			}
		}
	}
	return false
}

func (b board) checkWin() bool {
	// Check rows
	for r := range b {
		winner := true
		for c := range b[r] {
			if !b[r][c].called {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}

	// Check columns
	for c := range b[0] {
		winner := true
		for r := range b[c] {
			if !b[r][c].called {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}

	return false
}

func (b board) sumByMark(m bool) int {
	sum := 0
	for _, r := range b {
		for _, s := range r {
			if s.called == m {
				sum += s.number
			}
		}
	}
	return sum
}

func readBoards(input []string) []board {
	boards := []board{}

	// Skip draws and first blank line
	for i := 2; i < len(input)-4; i += 6 {
		board := board{}
		for _, row := range input[i : i+5] {
			boardRow := []boardSpace{}
			for _, space := range strings.Fields(row) {
				boardRow = append(boardRow, boardSpace{util.Atoi(space), false})
			}
			board = append(board, boardRow)
		}
		boards = append(boards, board)
	}

	return boards
}

func A(input []string) int {
	draws := util.ReadIntString(input[0])
	boards := readBoards(input)

	for _, draw := range draws {
		for _, b := range boards {
			b.mark(draw)
			if b.checkWin() {
				return draw * b.sumByMark(false)
			}
		}
	}

	return 0
}

func B(input []string) int {
	draws := util.ReadIntString(input[0])
	boards := readBoards(input)

	boardsLeft := len(boards)
	winners := make([]bool, len(boards))
	for _, draw := range draws {
		for i, b := range boards {
			if winners[i] {
				continue
			}
			b.mark(draw)
			if b.checkWin() {
				winners[i] = true
				boardsLeft--
			}
			if boardsLeft == 0 {
				return draw * b.sumByMark(false)
			}
		}
	}

	return 0
}
