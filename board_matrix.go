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

func (bm *BoardMatrix) CheckMove(move *Move, rack *Rack) (error, int) {
	if len(*move) == 0 {
		return errors.New("Empty move"), 0
	}
	err, orientation, index := move.GetOrientation()
	if err != nil {
		// Tiles are scattered on the board
		return err, 0
	}
	var words []*Tiles
	startx, starty := bm.getStart(move, orientation, index)

	words = append(words, bm.getWord(startx, starty, orientation, move, rack))


	//FIXME: continue from here
	panic(index)
	points := 0
	return nil, points
}

func (bm *BoardMatrix) getWord(x, y int, or Orientation, move *Move, rack *Rack) *Tiles {
	var tiles Tiles
	for {
		// Check board for letter
		err, tile := bm.GetTile(x, y)
		if tile != nil && err != nil {
			tiles = append(tiles, tile)
		} else {
			// Check move for letter
			err, tile := move.GetTile(x, y, rack)
			if tile != nil && err != nil {
				tiles = append(tiles, tile)
			} else {
				// Empty space || end of the board == terminator
				break;
			}
		}
		if or == OR_HORIZONTAL {
			x++
		} else {
			y++
		}
	}
	return &Tiles{}
}

func (bm *BoardMatrix) getStart(move *Move, or Orientation, index int) (int, int) {
	x := 0
	y := 0
	switch or {
	// Look up
	case OR_VERTICAL:
		x = index
		y = move.getMinY()
		for by := y; by > -1; by-- {
			_, tile := bm.GetTile(x, by)
			if tile == nil {
				y = by
				break
			}
		}
		y = 0
	// Look left
	case OR_HORIZONTAL:
		x = move.getMinX()
		y = index
		for bx := x; bx > -1; bx-- {
			_, tile := bm.GetTile(bx, y)
			if tile == nil {
				x = bx
				break
			}
		}
		x = 0

	}
	return x, y
}
