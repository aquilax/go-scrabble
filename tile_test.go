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

func TestFillRack(t *testing.T) {
	_, letters := NewLetters(LANG_EN)
	tiles := letters.NewTiles()
	rack := tiles.NewRack()
	_, _ = rack.PickTile(0)
	_, _ = rack.PickTile(1)
	rack_size := rack.size()
	tiles.FillRack(rack)
	if rack.size() != RACK_SIZE || rack_size >= rack.size() {
		t.Error("Rack not refilled properly")
	}
}

func TestDrawTile(t *testing.T) {
	_, letters := NewLetters(LANG_EN)
	tiles := letters.NewTiles()
	tiles_count := len(*tiles)
	var tile *Tile
	for i := 0; i <= tiles_count; i++ {
		tile = tiles.DrawTile()
	}
	if tile != nil {
		t.Error("Tiles are not removed properly")
	}
}

func TestPickTile(t *testing.T) {
	_, letters := NewLetters(LANG_EN)
	tiles := letters.NewTiles()
	rack := tiles.NewRack()
	err, _ := rack.PickTile(-1)
	if err == nil {
		t.Error("Wrong tile index check not working")
	}
	_, tile := rack.PickTile(0)
	if tile == nil {
		t.Error("Tile is nil")
	}
	err, tile2 := rack.PickTile(0)
	if err == nil || tile2 != nil {
		t.Error("Empty tile check not working")
	}
}

func Testsize(t *testing.T) {
	_, letters := NewLetters(LANG_EN)
	tiles := letters.NewTiles()
	rack := tiles.NewRack()
	if rack.size() != RACK_SIZE {
		t.Error("Wrong rack size")
	}
	_, _ = rack.PickTile(0)
	if rack.size() != RACK_SIZE-1 {
		t.Error("Wrong rack size")
	}
}
