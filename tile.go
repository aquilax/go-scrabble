package scrabble

type Tile struct {
	char   rune
	points int
}

// Creates new tile
func NewTile(char rune) *Tile {
	//TODO: Get actual tile points
	return &Tile{char, 0}
}
