package scrabble

const (
	// Size of the board
	BOARD_SIZE = 15

	// Plain place
	PLAIN  = 1
	// Doubles the letter score
	DBL_LETTER = iota
	// Tripples the letter score
	TRIPPLE_LETTER
	// Doubles the word score
	DOUBLE_WORD
	// Tripples the word score
	TRIPPLE_WORD
)


type Board [BOARD_SIZE][BOARD_SIZE]int

// Creates new Scrabble board
func NewBoard() *Board {
	return &Board{
		{5,1,1,2,1,1,1,5,1,1,1,2,1,1,5},
		{1,4,1,1,1,3,1,1,1,3,1,1,1,4,1},
		{1,1,4,1,1,1,2,1,2,1,1,1,4,1,1},
		{2,1,1,4,1,1,1,2,1,1,1,4,1,1,2},
		{1,1,1,1,4,1,1,1,1,1,4,1,1,1,1},
		{1,3,1,1,1,3,1,1,1,3,1,1,1,3,1},
		{1,1,2,1,1,1,2,1,2,1,1,1,2,1,1},
		{5,1,1,2,1,1,1,4,1,1,1,2,1,1,5},
		{1,1,2,1,1,1,2,1,2,1,1,1,2,1,1},
		{1,3,1,1,1,3,1,1,1,3,1,1,1,3,1},
		{1,1,1,1,4,1,1,1,1,1,4,1,1,1,1},
		{2,1,1,4,1,1,1,2,1,1,1,4,1,1,2},
		{1,1,4,1,1,1,2,1,2,1,1,1,4,1,1},
		{1,4,1,1,1,3,1,1,1,3,1,1,1,4,1},
		{5,1,1,2,1,1,1,5,1,1,1,2,1,1,5},
	}
}

// Gets the score multiplier for tile with coordinates x,y
func (b *Board) GetMultiplier(x, y int) (error, int) {
	err := checkBoundary(x,y)
	if err != nil {
		return err, 0
	}
	return nil, b[y][x]
}
