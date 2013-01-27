package scrabble

import (
	"errors"
	"math/rand"
)

type Tile struct {
	letter string
	points int
}

type Tiles []*Tile

type Rack [RACK_SIZE]*Tile

func (tl *Tiles) NewRack() *Rack {
	var rack Rack
	for i := 0; i < RACK_SIZE; i++ {
		rack[i] = tl.DrawTile()
	}
	return &rack
}

func (tl *Tiles) FillRack(rack *Rack) {
	for i := 0; i < RACK_SIZE; i++ {
		if rack[i] != nil {
			continue
		}
		rack[i] = tl.DrawTile()
		if rack[i] == nil {
			// No more tiles left
			break
		}
	}
}

func (tl *Tiles) DrawTile() *Tile {
	size := len(*tl)
	if size == 0 {
		return nil
	}
	index := 0
	if size > 1 {
		index = rand.Intn(size - 1)
	}
	tile := (*tl)[index]
	// Remove tile from list
	*tl = append((*tl)[:index], (*tl)[index+1:]...)
	return tile
}

func (r *Rack) PickTile(index int) (error, *Tile) {
	if index < 0 || index > RACK_SIZE {
		return errors.New("Incorrect tile index"), nil
	}
	tile := (*r)[index]
	if tile == nil {
		return errors.New("No tile on this place"), nil
	}
	(*r)[index] = nil
	return nil, tile
}

func (r *Rack) size() int {
	result := 0
	for i := 0; i < RACK_SIZE; i++ {
		if (*r)[i] != nil {
			result++
		}
	}
	return result
}
