package scrabble

import (
	"errors"
)

// Checks if the coordinates are withing the board limits
func checkBoundary (x, y int) error {
	if x >= BOARD_SIZE || y >= BOARD_SIZE {
		return errors.New("Coordinates out of board")
	}
	return nil
}
