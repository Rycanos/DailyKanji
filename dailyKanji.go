package main

import (
	"dailyKanji/character"
	"flag"
	"fmt"
	"time"
)

const TIME_USAGE string = "Sets the time for when the program should display the next randomly picked Kanji\n"
const TIME_DEFAULT string = "07:00"

const JLPT_USAGE string = "Sets the characters to display by their inclusion in the JLPT from level 3 to 5 (5 being the easiest)\n"
const JLPT_DEFAULT int = 5

const STROKES_USAGE string = "Sets the display of strokes for the given Kanji\n"
const STROKES_DEFAULT bool = false

const DATA_USAGE string = "Sets the path to the data file of the Kanji list\n"
const DATA_DEFAULT string = "Data/Kanji_20250717_140306.xml"

// TODO: Flags:
//  -time="HH:mm AM" customize the time when the program should display the character
//  -jlvl=[3-5] display only character from the JLPT of chosen lvl
//  -display-strokes display number of strokes for now (animation of strokes later)
//  -data="PATH" change the input file for the Kanji data

func displayCharacter(char character.Character) error {
	fmt.Printf("Kanji picked: %s\n", char.Char)
	// TODO: display with GUI
	return nil
}

func main() {
	// Getting time of launch
	startTime := time.Now()

	// Setting flags for the program
	timePtr := flag.String("time", TIME_DEFAULT, TIME_USAGE)
	jlptPtr := flag.Int("jlvl", JLPT_DEFAULT, JLPT_USAGE)
	strokesPtr := flag.Bool("display-strokes", STROKES_DEFAULT, STROKES_USAGE)
	dataPathPtr := flag.String("data", DATA_DEFAULT, DATA_USAGE)

	flag.Parse()

	fmt.Println("DEBUG FLAGS ------")
	fmt.Println("timePtr: ", *timePtr)
	fmt.Println("jlptPtr: ", *jlptPtr)
	fmt.Println("strokesPtr: ", *strokesPtr)
	fmt.Println("dataPathPtr: ", *dataPathPtr)
	fmt.Println("------------------")

	// Loading data set
	character.LoadCharactersFromSheet(*dataPathPtr)

	// Picking and displaying the first character
	char, err := character.PickCharacter()
	if err != nil {
		fmt.Println(err)
		return
	}
	displayCharacter(char)

	// Parsing the value of timePtr "15:04" corresponds to hours and minutes
	targetTime, errTParse := time.Parse("15:04", *timePtr)
	if errTParse != nil {
		fmt.Println(errTParse)
		return
	}

	//TODO Cleanup
	if false {
		fmt.Println("In if false")
		// Check the amount of time between start and the next programmed display at timePtr
		startTimeNextDay := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), targetTime.Hour(), targetTime.Minute(), 0, 0, startTime.Location())
		startTimeNextDay = startTimeNextDay.AddDate(0, 0, 1)
		diff := startTimeNextDay.UnixMilli() - startTime.UnixMilli()

		fmt.Println(diff)
		// Wait until next display
		//		time.Sleep(diff)
	}

	/* 	tickerDay := time.NewTicker(24 * time.Hour) */
	tickerDay := time.NewTicker(1 * time.Second)
	// Making channel to send a signal when all the characters have been cycled through
	Done := make(chan bool)

	// Main loop goroutine anonymous function
	go func() {
		for {
			select {
			// tickerDay.C fires at tick defined on construction of tickerDay and returns Time
			case t := <-tickerDay.C:
				// Picking a new character
				char, err = character.PickCharacter()
				// Checking if it is the last character
				if char == (character.Character{}) {
					Done <- true
					return
				}
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Picked new character at: ", t)
				displayCharacter(char)
			}
		}
	}()
	<-Done
	tickerDay.Stop()
	fmt.Println("Ticker stopped at: ", time.Since(startTime))
}
