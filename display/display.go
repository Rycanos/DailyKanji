package display

import (
	"dailyKanji/character"
	"fmt"
)

func DisplayCharacter(char character.Character) error {
	fmt.Printf("Kanji picked: %s   jlpt: %d   id: %d\n", char.Char, char.JlptLvl, char.CharId)
	// TODO: display with GUI
	return nil
}
