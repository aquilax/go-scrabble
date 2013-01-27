package scrabble

import (
	"errors"
)

type BoardMatrix [BOARD_SIZE][BOARD_SIZE]*Tile

// Creates new board matrix
func NewBoardMatrix() *BoardMatrix {
	return &BoardMatrix{}
}

// Cleans the board matrix
func (bm *BoardMatrix) Clean() {
	for y := 0; y < BOARD_SIZE; y++ {
		for x := 0; x < BOARD_SIZE; x++ {
			bm[y][x] = nil
		}
	}
}

// Get tile at coordinates x, y
func (bm *BoardMatrix) GetTile(x, y int) (error, *Tile) {
	err := checkBoundary(x, y)
	if err != nil {
		return err, nil
	}
	return nil, bm[y][x]
}

// Place tile at coordinates x, y
func (bm *BoardMatrix) SetTile(x, y int, tile *Tile) error {
	err := checkBoundary(x, y)
	if err != nil {
		return err
	}
	if bm[y][x] != nil {
		return errors.New("Tile is already placed on this position")
	}
	bm[y][x] = tile
	return nil
}

func (bm *BoardMatrix) IsFull() bool {
	for y := 0; y < BOARD_SIZE; y++ {
		for x := 0; x < BOARD_SIZE; x++ {
			if bm[y][x] != nil {
				return false
			}
		}
	}
	return true
}

func (bm *BoardMatrix) CheckMove(move *Move) (error, int) {
	if len(*move) == 0 {
		return errors.New("Empty move"), 0
	}
	err, direction := move.GetOrientation()
	if err != nil {
		return err, 0
	}
	//FIXME: continue from here
	panic(direction)
	points := 0
	return nil, points
}
