package scrabble

import "testing"

func TestcheckBoundary(t *testing.T) {
	err := checkBoundary(0, 0)
	if err != nil {
		t.Error("Wrong board size")
	}
	err1 := checkBoundary(BOARD_SIZE, 0)
	if err1 == nil {
		t.Error("Wrong boundary check")
	}
	err2 := checkBoundary(0, BOARD_SIZE)
	if err2 == nil {
		t.Error("Wrong boundary check")
	}
	err3 := checkBoundary(-1, -1)
	if err3 == nil {
		t.Error("Wrong boundary check")
	}
}
