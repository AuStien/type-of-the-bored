package main

import (
	"flag"
	"fmt"

	"github.com/eiannone/keyboard"

	"github.com/austien/type-of-the-bored/ansi"
	"github.com/austien/type-of-the-bored/terminal"
	"github.com/austien/type-of-the-bored/words"
)

var triggerFetchWord = make(chan bool)
var nextWord = make(chan *words.Text)

func main() {

	useWordFlag := flag.Bool("word", false, "Type words instead of quotes")

	flag.Parse()

	cols, err := terminal.GetCols()
	if err != nil {
		panic(err)
	}

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()

		// Capture panics, to show cursor
		err := recover()
		if err != nil {
			fmt.Print(ansi.SHOW_CURSOR)
			fmt.Printf("\nerr: %s\n", err)
		}
	}()

	fmt.Printf("Press ESC or Ctrl+C to quit\n\n")

	// Fetch next word async
	go func() {
		for {
			<-triggerFetchWord
			w, err := words.NewText(*useWordFlag, cols)
			if err != nil {
				panic(err)
			}
			nextWord <- w
		}
	}()

	word, err := words.NewText(*useWordFlag, cols)
	if err != nil {
		panic(err)
	}

	triggerFetchWord <- true

	// Hide cursor
	fmt.Print(ansi.HIDE_CURSOR)

	for {
		fmt.Printf("%s%s%s%s", ansi.CLEAR_LINE, ansi.MOVE_CURSOR_UP, ansi.CLEAR_LINE, word.ToString())

		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		// Break if ESC or Ctrl + C is inputed
		if event.Key == keyboard.KeyEsc || event.Key == keyboard.KeyCtrlC {
			break
		}

		// Handle backspace
		if event.Key == keyboard.KeyBackspace || event.Key == keyboard.KeyBackspace2 {
			word.Previous()
			continue
		}

		word.CompareCurrentLetter(event.Rune)

		if word.IsComplete() {
			word = <-nextWord
			triggerFetchWord <- true
		} else {
			word.Next()
		}
	}

	// 	Show cursor
	fmt.Printf("%s\n", ansi.SHOW_CURSOR)
}
