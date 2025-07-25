package display

import (
	"dailyKanji/character"
	"fmt"
)

func DisplayCharacter(char character.Character) error {
	fmt.Printf("Kanji picked: %s\n", char.Char)
	// TODO: display with GUI
	return nil
}
