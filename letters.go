package scrabble

import (
	"errors"
)

// Letter distribution and tile points from:
// http://en.wikipedia.org/wiki/Scrabble_letter_distributions

type Letter struct {
	count  int
	points int
}

// Represents set of letter rules
// String is used as a key because in some languages
// a tile has more than one rune
type Letters map[string]Letter

func NewLetters(language string) (error, *Letters) {
	switch language {
	case "en":
		return nil, &Letters{
			" ": {2, 0},
			"E": {12, 1},
			"A": {9, 1},
			"I": {9, 1},
			"O": {8, 1},
			"N": {6, 1},
			"R": {6, 1},
			"T": {6, 1},
			"L": {4, 1},
			"S": {4, 1},
			"U": {4, 1},
			"D": {4, 2},
			"G": {3, 2},
			"B": {2, 3},
			"C": {2, 3},
			"M": {2, 3},
			"P": {2, 3},
			"F": {2, 4},
			"H": {2, 4},
			"V": {2, 4},
			"W": {2, 4},
			"Y": {2, 4},
			"K": {1, 5},
			"J": {1, 8},
			"X": {1, 8},
			"Q": {1, 10},
			"Z": {1, 10},
		}
	}
	return errors.New("Language not defined"), nil
}
