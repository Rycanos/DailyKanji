package character

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

const charSheetFile string = "Data/Kanji_20250717_140306.xml"

type Character struct {
	//	XMLName     xml.Name `xml:"Kanji_data"`
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

func LoadCharactersFromSheet(charSheet string) {
	if charSheet == "" {
		charSheet = charSheetFile
	}
	xmlFile, err := os.Open(charSheet)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully opened \"%s\"\n", charSheet)

	byteValue, _ := io.ReadAll(xmlFile)
	var characters Characters

	xml.Unmarshal(byteValue, &characters)

	fmt.Printf("\n----PRINTING CONTENTS----\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("\n%s\tid: %d, strokes: %d", characters.CharList[i].Char, characters.CharList[i].CharId, characters.CharList[i].CharStroke)
	}

	defer xmlFile.Close()
}
