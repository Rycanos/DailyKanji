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

func LoadCharactersFromSheet(charSheet string) error {
	if charSheet == "" {
		return errors.New("Path to character sheet is invalid")
	}
	xmlFile, err := os.Open(charSheet)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully opened \"%s\"\n", charSheet)

	byteValue, _ := io.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &characters)

	// Initializing
	var length int = len(characters.CharList)
	if length <= 0 {
		return errors.New("Character list is empty")
	}

	// Filling a list of int with
	for i := 0; i < length; i++ {
		randListInt = append(randListInt, i)
	}

	defer xmlFile.Close()
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
