package scrabble

import (
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

func (tl *Tiles) DrawTile() *Tile {
	size := len(*tl)
	index := rand.Intn(size - 1)
	tile := (*tl)[index]
	// Remove tile from list
	*tl = append((*tl)[:index], (*tl)[index+1:]...)
	return tile
}
