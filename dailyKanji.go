package main

import (
	"DailyKanji/character"
	"context"
	"embed"
	"flag"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

const TIME_USAGE string = "Sets the time for when the program should display the next randomly picked Kanji\n"
const TIME_DEFAULT string = "07:00"

const JLPT_USAGE string = "Sets the characters to display by their inclusion in the JLPT from level 3 to 5 (5 being the easiest)\n"
const JLPT_DEFAULT int = 5

const STROKES_USAGE string = "Sets the display of strokes for the given Kanji\n"
const STROKES_DEFAULT bool = false

const DATA_USAGE string = "Sets the path to the data file of the Kanji list\n"
const DATA_DEFAULT string = "Data/Kanji_20250717_140306.xml"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

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
	if character.LoadCharactersFromSheet(*dataPathPtr, *jlptPtr) != nil {
		return
	}

	// Picking and displaying the first character
	char, err := character.PickCharacter()
	if err != nil {
		fmt.Println(err)
		return
	}
	a.DisplayCharacter(char)

	// Parsing the value of timePtr "15:04" corresponds to hours and minutes
	targetTime, errTParse := time.Parse("15:04", *timePtr)
	if errTParse != nil {
		fmt.Println(errTParse)
		return
	}

	// Calculating the time for next display and returning the ticker
	// TODO: custom ticker interval (as parameter?)
	tickerDay := calculateTicker(startTime, targetTime)

	// Making channel to send a signal when all the characters have been cycled through
	Done := make(chan bool)

	// Main loop goroutine call, displays a character at each tick
	go a.manageTicker(tickerDay, Done)

	<-Done
	tickerDay.Stop()
	fmt.Println("The program exited after: ", time.Since(startTime))

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// TODO: Character picker for frontend button
/* func (a *App) PickCharacter() (character.Character, error) {
	// Call the character package to pick a character
	char, err := character.PickCharacter()
	if err != nil {
		return character.Character{}, err
	}
	// Display the character using the display package
	err = display.DisplayCharacter(char)
	if err != nil {
		return character.Character{}, err
	}
	return char, nil
} */

func (a *App) DisplayCharacter(char character.Character) error {
	fmt.Printf("Kanji picked: %s   id: %d\n", char.Char, char.CharId)

	runtime.EventsEmit(a.ctx, "characterPicked", char)
	return nil
}

// TODO: Flags:
//  -time="HH:mm AM" customize the time when the program should display the character
//  -jlvl=[3-5] display only character from the JLPT of chosen lvl
//  -display-strokes display number of strokes for now (animation of strokes later)
//  -data="PATH" change the input file for the Kanji data

func (a *App) manageTicker(tickerDay *time.Ticker, Done chan<- bool) {
	for {
		// tickerDay.C fires at tick defined on construction of tickerDay and returns Time
		<-tickerDay.C
		// Picking a new character
		char, err := character.PickCharacter()
		// Checking if it is the last character
		if char == (character.Character{}) {
			Done <- true
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		a.DisplayCharacter(char)
	}
}

func calculateTicker(startTime time.Time, targetTime time.Time) (ticker *time.Ticker) {
	// Check the amount of time between start and the next programmed display at timePtr
	/* 	startTimeNextDay := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), targetTime.Hour(),
	targetTime.Minute(), 0, 0, startTime.Location()) */
	// Uncomment for debug purposes
	startTimeNextDay := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(),
		startTime.Minute(), startTime.Second()+5, 0, startTime.Location())

	/* 	startTimeNextDay = startTimeNextDay.AddDate(0, 0, 1) */
	diff := startTimeNextDay.UnixNano() - startTime.UnixNano()

	fmt.Println("time.Duration(diff): ", time.Duration(diff))
	// Wait until next display
	time.Sleep(time.Duration(diff))

	// Sets the display of characters to be every 24 hours
	/* 	return time.NewTicker(24 * time.Hour) */
	// Uncomment for debug purposes
	return time.NewTicker(time.Second / 100)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	errWails := wails.Run(&options.App{
		Title:  "DailyKanji",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if errWails != nil {
		println("Error:", errWails.Error())
	}

}
