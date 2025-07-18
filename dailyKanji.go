package main

import (
	"dailyKanji/character"
	"fmt"
)

func main() {
	fmt.Printf("DailyKanji displays a new Kanji character every day to help you learn!\n")
	character.LoadCharactersFromSheet("")
	char, err := character.PickCharacter()
	fmt.Printf("%s\n", err)
	fmt.Printf("first Kanji: %s\n", char.Char)
}
