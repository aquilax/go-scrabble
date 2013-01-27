package scrabble

import (
	"errors"
)

const (
	MAX_PLAYERS = 4
)

type Player struct {
	id   int
	rack *Rack
}

type Players []*Player

func NewPlayers() *Players {
	return &Players{}
}

func (pl *Players) NewPlayer(tiles *Tiles) (error, *Player) {
	if len(*pl) == MAX_PLAYERS {
		return errors.New("No more players allowed"), nil
	}
	var player Player
	*pl = append(*pl, &player)
	player.id = len(*pl) - 1
	player.rack = tiles.NewRack()
	return nil, &player
}
