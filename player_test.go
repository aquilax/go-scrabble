package scrabble

import "testing"

func TestNewPlayers(t *testing.T) {
	players := NewPlayers()
	if players == nil {
		t.Error("Players is nil")
	}
}

func TestNewPlayer(t *testing.T) {
	_, letters := NewLetters(LANG_EN)
	tiles := letters.NewTiles()
	players := NewPlayers()
	_, player := players.NewPlayer(tiles)
	if player == nil {
		t.Error("Player is nil")
	}
	var err error
	for i := 0; i < MAX_PLAYERS; i++ {
		err, _ = players.NewPlayer(tiles)
	}
	if err == nil {
		t.Error("MAX_PLAYERS check faled")
	}
}
