package scrabble

import (
	"errors"
)

const (
	OR_HORIZONTAL = true
	OR_VERTICAL   = false
)

type Orientation bool

type MoveTile struct {
	x int
	y int
}

type Move map[int]*MoveTile

func NewMove() *Move {
	return &Move{}
}

func (m *Move) AddMove(rack_id, x, y int) error {
	err := checkBoundary(x, y)
	if err != nil {
		return err
	}
	(*m)[rack_id] = &MoveTile{x, y}
	return nil
}

func (m *Move) Clean() {
	m = NewMove()
}

func (m *Move) GetOrientation() (error, Orientation) {
	last_x := -1
	last_y := -1
	is_horizontal := true
	is_vertical := true
	for _, move_tile := range *m {
		if last_x != -1 && last_x != move_tile.x {
			is_vertical = false
		}
		last_x = move_tile.x

		if last_y != -1 && last_y != move_tile.y {
			is_horizontal = false
		}
		last_y = move_tile.y
	}
	if is_horizontal {
		return nil, OR_HORIZONTAL
	}
	if is_vertical {
		return nil, OR_VERTICAL
	}
	return errors.New("Bad move: neither horizontal not vertical"), false
}
