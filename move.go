package scrabble

import (
	"errors"
)

const (
	OR_HORIZONTAL = true
	OR_VERTICAL   = false

	MAXINT = 1000
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

func (m *Move) GetOrientation() (error, Orientation, int) {
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
		return nil, OR_HORIZONTAL, last_y
	}
	if is_vertical {
		return nil, OR_VERTICAL, last_x
	}
	return errors.New("Bad move: neither horizontal not vertical"), false, -1
}

func (m *Move) getMinX() int {
	minx := MAXINT
	for _, tile := range *m {
		if tile.x < minx {
			minx = tile.x
		}
	}
	return minx
}

func (m *Move) getMinY() int {
	miny := MAXINT
	for _, tile := range *m {
		if tile.y < miny {
			miny = tile.y
		}
	}
	return miny
}

func (m *Move) GetTile(x, y int, rack *Rack) (error, *Tile) {
	for rack_id, mt := range *m {
		if mt.x == x && mt.y == y {
			return nil, rack[rack_id]
		}
	}
	return errors.New("Tile not found on this position"), nil
}
