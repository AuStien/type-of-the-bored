package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

type Fetch struct {
	Words []Word
}

type Word struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
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
				fmt.Printf("\033[1A\033[K\033[1A\033[K")
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
	resp, err := http.Get("https://randomword.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var word Word
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "random_word" {
					word.Word = n.FirstChild.Data
				} else if a.Key == "id" && a.Val == "random_word_definition" {
					word.Definition = n.FirstChild.Data
				}
			}
		}

		// traverses the HTML of the webpage from the first child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return &word, nil
}

func WordStr(w *Word, expert bool) string {
	if expert {
		return fmt.Sprintf("%s\n", w.Definition)
	} else {
		return fmt.Sprintf("%s - %s\n", w.Word, w.Definition)
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
