package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

type Fetch struct {
	Words []Word
}

type Word struct {
	Word          string `json:"word"`
	Definition    string `json:"definition"`
	Pronunciation string `json:"pronunciation"`
}

var totbCmd = &cobra.Command{
	Use: "totb",
	Run: func(cmd *cobra.Command, args []string) {
		word, err := GetWord()
		if err != nil {
			panic(err)
		}

		expert, err := cmd.Flags().GetBool("expert")
		if err != nil {
			panic(err)
		}

		fmt.Print(WordStr(word, expert))
    fmt.Print("\n\033[1A")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == word.Word {
				fmt.Printf("\033[1A\033[K\033[1A\033[K\u2714 %s", WordStr(word, false))
				word, err = GetWord()
				if err != nil {
					panic(err)
				}

				fmt.Print(WordStr(word, expert))
			} else {
				fmt.Print("\033[1A\033[K")
			}
      fmt.Print("\n\033[1A")
		}
		if scanner.Err() != nil {
			panic(scanner.Err)
		}

	},
}

func GetWord() (*Word, error) {
	resp, err := http.Get("https://random-words-api.vercel.app/word")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var words []Word
	if err := json.NewDecoder(resp.Body).Decode(&words); err != nil {
		return nil, err
	}
	return &words[0], nil
}

func WordStr(w *Word, expert bool) string {
	if expert {
		return fmt.Sprintf("%s - %s\n", w.Pronunciation, w.Definition)
	} else {
		return fmt.Sprintf("%s (%s) - %s\n", w.Word, w.Pronunciation, w.Definition)
	}
}

func init() {
	totbCmd.Flags().BoolP("expert", "e", false, "Expert Mode")
}

func main() {
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan

		fmt.Println("")

		os.Exit(0)
	}()
	totbCmd.Execute()
}
