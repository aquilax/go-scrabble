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
	_, orientationv := move.GetOrientation()
	if orientationv != OR_VERTICAL {
		t.Error("Wrong orientation detected")
	}
	move.Clean()
	move.AddMove(0, 0, 0)
	move.AddMove(1, 1, 0)
	move.AddMove(2, 2, 0)
	_, orientationh := move.GetOrientation()
	if orientationh != OR_HORIZONTAL {
		t.Error("Wrong orientation detected")
	}
	move.Clean()
	move.AddMove(0, 0, 1)
	move.AddMove(1, 1, 0)
	err, _ := move.GetOrientation()
	if err == nil {
		t.Error("Wrong orientation detected")
	}
}
