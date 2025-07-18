package character

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
)

const charSheetFile string = "Data/Kanji_20250717_140306.xml"

type Character struct {
	CharId      int    `xml:"id"`
	CharStroke  int    `xml:"Strokes"`
	JlptLvl     int    `xml:"JLPT-test"`
	Char        string `xml:"Kanji"`
	ReadingJoyo string `xml:"Reading_within_Joyo"`
	MeaningOn   string `xml:"Translation_of_On"`
	MeaningKun  string `xml:"Translation_of_Kun"`
	Viewed      bool
}

type Characters struct {
	XMLName  xml.Name    `xml:"database"`
	CharList []Character `xml:"Kanji_data"`
}

var characters Characters
var randListInt []int

func LoadCharactersFromSheet(charSheet string) error {
	if charSheet == "" {
		charSheet = charSheetFile
	}
	xmlFile, err := os.Open(charSheet)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully opened \"%s\"\n", charSheet)

	byteValue, _ := io.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &characters)

	// fmt.Printf("\n----PRINTING CONTENTS----\n")
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%s\tid: %d, strokes: %d\n", characters.CharList[i].Char, characters.CharList[i].CharId, characters.CharList[i].CharStroke)
	// }

	// Initializing
	var length int = len(characters.CharList)
	if length <= 0 {
		return errors.New("Character list is empty")
	}

	// filling a list of int with
	for i := 0; i < length; i++ {
		randListInt = append(randListInt, i)
	}

	defer xmlFile.Close()
	return nil
}

func PickCharacter() (Character, error) {
	r := rand.Intn(len(randListInt))

	//TODO: pick a character at random using r, remove it from the list to shorten it
	// need a infinite loop in main waiting for an input of the user or that 24h has elapsed
	return characters.CharList[r], nil
}
