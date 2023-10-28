package words

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/austien/type-of-the-bored/ansi"
	"golang.org/x/net/html"
)

//go:embed quotes.json
var quotes embed.FS

type Text struct {
	Letters         []letter
	Word            string
	Description     string
	currentLetterNo int
}

func NewText(useWord bool, maxLength int) (*Text, error) {
	var fetchedWord string
	var description string
	var err error
	if useWord {
		fetchedWord, description, err = fetchWord()
	} else {
		fetchedWord, description, err = fetchQuote(maxLength)
	}
	if err != nil {
		return nil, err
	}

	w := &Text{
		Word:            fetchedWord,
		Description:     description,
		Letters:         generateLetters(fetchedWord),
		currentLetterNo: 0,
	}

	return w, nil
}

func (w *Text) ToString() string {
	str := ""

	for _, l := range w.Letters {
		str = fmt.Sprintf("%s%s", str, l.ToString())
	}

	str = fmt.Sprintf("%s\n- %s", str, w.Description)

	return str
}

func (w *Text) CompareCurrentLetter(input rune) {
	w.Letters[w.currentLetterNo].CompareCharacter(input)
}

func (w *Text) Next() {
	if w.currentLetterNo < len(w.Letters)-1 {
		w.Letters[w.currentLetterNo].HasUnderline = false
		w.currentLetterNo++
		w.Letters[w.currentLetterNo].HasUnderline = true
	}
}

func (w *Text) Previous() {
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

func (w *Text) IsComplete() bool {
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

func fetchQuote(maxLength int) (string, string, error) {
	b, err := quotes.ReadFile("quotes.json")
	if err != nil {
		return "", "", err
	}

	type Quote struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	r := []Quote{}
	if err := json.Unmarshal(b, &r); err != nil {
		return "", "", err
	}

	rand.New(rand.NewSource(time.Now().Unix()))

	if maxLength <= 0 {
		q := r[rand.Intn(len(r))]
		return q.Content, q.Author, nil
	}

	MAX_RETRIES := 100
	for i := 0; i < MAX_RETRIES; i++ {
		q := r[rand.Intn(len(r))]
		if len(q.Content) <= maxLength {
			return q.Content, q.Author, nil
		}
	}

	return "", "", fmt.Errorf("could not find quote with length less than %d", maxLength)
}
