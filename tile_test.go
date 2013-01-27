package scrabble

import "testing"

func TestNewRack(t *testing.T) {
	_, letters := NewLetters(LANG_EN)
	tiles := letters.NewTiles()
	old_tiles_count := len(*tiles)
	rack := tiles.NewRack()
	if rack == nil {
		t.Error("Rack is nil")
	}
	expected := old_tiles_count - RACK_SIZE
	actual := len(*tiles)
	if actual != expected {
		t.Errorf("Wrong number of tiles left actual:%d expected:%d", actual, expected)
	}
}
