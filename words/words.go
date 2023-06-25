package words

import (
	"fmt"
	"net/http"

	"github.com/austien/type-of-the-bored/ansi"
	"golang.org/x/net/html"
)

type Word struct {
	Letters         []letter
	Word            string
	Definition      string
	currentLetterNo int
}

func NewWord() (*Word, error) {
	fetchedWord, definition, err := fetchWord()
	if err != nil {
		return nil, err
	}

	w := &Word{
		Word:            fetchedWord,
		Definition:      definition,
		Letters:         generateLetters(fetchedWord),
		currentLetterNo: 0,
	}

	return w, nil
}

func (w *Word) ToString() string {
	str := ""

	for _, l := range w.Letters {
		str = fmt.Sprintf("%s%s", str, l.ToString())
	}

	str = fmt.Sprintf("%s - %s", str, w.Definition)

	return str
}

func (w *Word) CompareCurrentLetter(input rune) {
	w.Letters[w.currentLetterNo].CompareCharacter(input)
}

func (w *Word) Next() {
	if w.currentLetterNo < len(w.Letters)-1 {
		w.Letters[w.currentLetterNo].HasUnderline = false
		w.currentLetterNo++
		w.Letters[w.currentLetterNo].HasUnderline = true
	}
}

func (w *Word) Previous() {
	if w.currentLetterNo > 0 {
		w.Letters[w.currentLetterNo].HasUnderline = false
		w.Letters[w.currentLetterNo].Color = ansi.NONE
		w.currentLetterNo--
		w.Letters[w.currentLetterNo].HasUnderline = true
		w.Letters[w.currentLetterNo].Color = ansi.NONE
	} else {
		w.Letters[w.currentLetterNo].Color = ansi.NONE
	}
}

func (w *Word) IsComplete() bool {
	for _, l := range w.Letters {
		if l.Color.Name == ansi.RED.Name || l.Color.Name == ansi.NONE.Name {
			return false
		}
	}

	return true
}

func generateLetters(word string) []letter {
	letters := []letter{}

	for i, l := range word {
		if i == 0 {
			letters = append(letters, letter{Character: l, HasUnderline: true, Color: ansi.NONE})
		} else {
			letters = append(letters, letter{Character: l, HasUnderline: false, Color: ansi.NONE})
		}
	}
	return letters
}

func fetchWord() (string, string, error) {
	resp, err := http.Get("https://randomword.com")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", "", err
	}

	word := ""
	definition := ""

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "random_word" {
					word = n.FirstChild.Data
				} else if a.Key == "id" && a.Val == "random_word_definition" {
					definition = n.FirstChild.Data
				}
			}
		}

		// traverses the HTML of the webpage from the first child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return word, definition, nil
}
