package scrabble

import "testing"

var board_matrix *BoardMatrix
var letters *Letters

func TestNewBoardMatrix(t *testing.T) {
	board_matrix = NewBoardMatrix()
	if board_matrix == nil {
		t.Error("Board is nil")
	}
	_, letters = NewLetters(LANG_EN)
}

func TestClean(t *testing.T) {
	tile := letters.NewTile("A")
	_ = board_matrix.SetTile(0, 0, tile)
	board_matrix.Clean()
	_, get_tile := board_matrix.GetTile(0, 0)
	if get_tile != nil {
		t.Error("Clean board matrix failed")
	}
}

func TestGetTile(t *testing.T) {
	board_matrix.Clean()
	_, empty_tile := board_matrix.GetTile(0, 0)
	if empty_tile != nil {
		t.Error("Empty tile should be nil")
	}
	tile := letters.NewTile("A")
	_ = board_matrix.SetTile(0, 0, tile)
	err, get_tile := board_matrix.GetTile(0, 0)
	if err != nil {
		t.Error(err)
	}
	if get_tile != tile {
		t.Error("Got wrong tile")
	}
}

func TestSetTile(t *testing.T) {
	board_matrix.Clean()
	tile1 := letters.NewTile("A")
	_ = board_matrix.SetTile(0, 0, tile1)
	tile2 := letters.NewTile("A")
	err := board_matrix.SetTile(0, 0, tile2)
	if err == nil {
		t.Error("Tiles should not be placed one over another")
	}
}
