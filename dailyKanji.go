package main

import (
	"dailyKanji/character"
	"fmt"
)

func main() {
	fmt.Printf("DailyKanji displays a new Kanji character every day to help you learn!\n")
	character.LoadCharactersFromSheet("")
}
