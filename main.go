package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Fetch struct {
	Words []Word
}

type Word struct {
	Word           string `json:"word"`
	Definition     string `json:"definition"`
	Pronounciation string `json:"pronounciation"`
}

var totbCmd = &cobra.Command{
	Use: "totb",
	Run: func(cmd *cobra.Command, args []string) {

		word, err := GetWord()
		if err != nil {
			panic(err)
		}

		fmt.Println(word.Word)
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

func main() {
	totbCmd.Execute()
}
