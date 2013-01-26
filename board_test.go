package scrabble

import "testing"

var board *Board

func TestNew(t *testing.T) {
	board = NewBoard()
	if board == nil {
		t.Error("Board is nil")
	}
}

func TestGetMultiplier(t *testing.T) {
	var multiplier int
	_, multiplier = board.GetMultiplier(0, 0)
	if multiplier != TRIPPLE_WORD {
		t.Error("Wrong board")
	}
	_, multiplier = board.GetMultiplier(7, 7)
	if multiplier != DOUBLE_WORD {
		t.Error("Wrong board")
	}
	e, _ := board.GetMultiplier(BOARD_SIZE, 0)
	if e == nil {
		t.Error("Wrong error handling")
	}
}
