package scrabble

import "testing"

func TestNewLetters(t *testing.T) {
	err, _ := NewLetters(-1)
	if err == nil {
		t.Error("Error in error handling")
	}
	_, letters := NewLetters(LANG_EN)
	if (*letters)["A"].count != 9 || (*letters)["A"].points != 1 {
		t.Error("Wrong letter A")
		if (*letters)[" "].count != 2 || (*letters)[" "].points != 0 {
			t.Error("Wrong joker")
		}
	}
}
