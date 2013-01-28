package scrabble

import (
	"testing"
)

func TestNewMove(t *testing.T) {
	move := NewMove()
	if move == nil {
		t.Error("Move is nil")
	}
}

func TestAddMove(t *testing.T) {
	move := NewMove()
	move.AddMove(0, 1, 1)
	if len(*move) != 1 {
		t.Error("Wrong move length")
	}
	move.AddMove(0, 1, 2)
	if len(*move) != 1 {
		t.Error("Wrong move length")
	}
}

func TestGetOrientation(t *testing.T) {
	move := NewMove()
	move.AddMove(0, 0, 0)
	move.AddMove(1, 0, 1)
	move.AddMove(2, 0, 2)
	_, orientationv, indexx := move.GetOrientation()
	if indexx != 0 {
		t.Error("Wrong index detected")
	}
	if orientationv != OR_VERTICAL {
		t.Error("Wrong orientation detected")
	}
	move.Clean()
	move.AddMove(0, 0, 0)
	move.AddMove(1, 1, 0)
	move.AddMove(2, 2, 0)
	_, orientationh, indexy := move.GetOrientation()
	if indexy != 0 {
		t.Error("Wrong index detected")
	}
	if orientationh != OR_HORIZONTAL {
		t.Error("Wrong orientation detected")
	}
	move.Clean()
	move.AddMove(0, 0, 1)
	move.AddMove(1, 1, 0)
	err, _, _ := move.GetOrientation()
	if err == nil {
		t.Error("Wrong orientation detected")
	}
}

func TestgetMinX(t *testing.T) {
	move := NewMove()
	move.AddMove(0, 1, 0)
	move.AddMove(1, 2, 0)
	move.AddMove(2, 3, 0)
	x := move.getMinX()
	if x != 1 {
		t.Error("Wrong min x")
	}
}

func TestgetMinY(t *testing.T) {
	move := NewMove()
	move.AddMove(0, 0, 1)
	move.AddMove(1, 0, 2)
	move.AddMove(2, 0, 3)
	y := move.getMinY()
	if y != 1 {
		t.Error("Wrong min y")
	}
}
