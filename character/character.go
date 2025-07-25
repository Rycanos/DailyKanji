package character

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Character struct {
	CharId      int    `xml:"id"`
	CharStroke  int    `xml:"Strokes"`
	JlptLvl     int    `xml:"JLPT-test"`
	Char        string `xml:"Kanji"`
	ReadingJoyo string `xml:"Reading_within_Joyo"`
	MeaningOn   string `xml:"Translation_of_On"`
	MeaningKun  string `xml:"Translation_of_Kun"`
}

type Characters struct {
	XMLName  xml.Name    `xml:"database"`
	CharList []Character `xml:"Kanji_data"`
}

var characters Characters
var randListInt []int

func filterOutChars(filterJlpt int) {
	var temp Characters
	for i := range len(characters.CharList) {
		if characters.CharList[i].JlptLvl >= filterJlpt {
			/* 			fmt.Printf("Kanji filtered: %s   jlpt: %d   id: %d\n", characters.CharList[i].Char, characters.CharList[i].JlptLvl, characters.CharList[i].CharId) */
			temp.CharList = append(temp.CharList, characters.CharList[i])
		}
	}
	if len(temp.CharList) > 0 {
		characters.CharList = temp.CharList
	}
}

func LoadCharactersFromSheet(charSheet string, filterJlpt int) error {
	if charSheet == "" {
		return errors.New("path to character sheet is invalid")
	}
	xmlFile, err := os.Open(charSheet)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &characters)

	filterOutChars(filterJlpt)

	// Initializing
	var length int = len(characters.CharList)
	if length <= 0 {
		return errors.New("character list is empty")
	}

	// Filling a list of int with its index
	for i := 0; i < length; i++ {
		randListInt = append(randListInt, i)
	}

	return nil
}

func PickCharacter() (Character, error) {
	// returning a blank Character if the list is empty
	if len(randListInt) <= 0 {
		fmt.Println("Character list is empty!")
		return Character{}, nil
	}
	// Getting a random index
	r := rand.Intn(len(randListInt))
	// Getting the value at the index
	x := randListInt[r]
	// Cutting it from the list
	randListInt = append(randListInt[:r], randListInt[r+1:]...)

	// Returning the character at the random index
	return characters.CharList[x], nil
}
