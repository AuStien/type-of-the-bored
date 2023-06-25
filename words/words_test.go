package words

import (
	"errors"
	"fmt"
	"testing"

	"github.com/austien/type-of-the-bored/ansi"
)

func TestNewWord(t *testing.T) {
	word, err := NewWord()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(word.ToString())
}

func TestWord(t *testing.T) {
	const WORD = "foobar"

	word := Word{
		Word:       WORD,
		Letters:    generateLetters(WORD),
		Definition: "Lorem Ipsum",
	}

	if err := testLetter(word.Letters[word.currentLetterNo], 'f', ansi.NONE, true); err != nil {
		t.Fatal(err)
	}

	word.CompareCurrentLetter('f')
	if err := testLetter(word.Letters[word.currentLetterNo], 'f', ansi.GREEN, true); err != nil {
		t.Fatal(err)
	}

	word.Next()

	if err := compareHasUnderline(word.Letters[word.currentLetterNo-1].HasUnderline, false); err != nil {
		t.Fatal(err)
	}

	if err := testLetter(word.Letters[word.currentLetterNo], 'o', ansi.NONE, true); err != nil {
		t.Fatal(err)
	}

	word.CompareCurrentLetter('o')
	if err := testLetter(word.Letters[word.currentLetterNo], 'o', ansi.GREEN, true); err != nil {
		t.Fatal(err)
	}

	word.Next()

	if err := testLetter(word.Letters[word.currentLetterNo], 'o', ansi.NONE, true); err != nil {
		t.Fatal(err)
	}

	word.CompareCurrentLetter('x')
	if err := testLetter(word.Letters[word.currentLetterNo], 'o', ansi.RED, true); err != nil {
		t.Fatal(err)
	}
}

func compareLetters(expect, is rune) error {
	if expect != is {
		return fmt.Errorf("comparing letters: expected %s, was %s", string(expect), string(is))
	}
	return nil
}

func compareColors(expect, is ansi.CODE) error {
	if expect != is {
		return fmt.Errorf("comparing colors: expected %v, was %v", expect.Name, is.Name)
	}
	return nil
}

func compareHasUnderline(expect, is bool) error {
	if expect != is {
		return fmt.Errorf("comparing underline: expected %t, was %t", expect, is)
	}
	return nil
}

func testLetter(letter letter, char rune, color ansi.CODE, hasUnderline bool) error {
	errorMessage := ""

	if err := compareLetters(letter.Character, char); err != nil {
		errorMessage = fmt.Sprintf("%s\n", err)
	}

	if err := compareColors(letter.Color, color); err != nil {
		errorMessage = fmt.Sprintf("%s%s\n", errorMessage, err)
	}

	if err := compareHasUnderline(letter.HasUnderline, hasUnderline); err != nil {
		errorMessage = fmt.Sprintf("%s%s\n", errorMessage, err)
	}

	if errorMessage == "" {
		return nil
	}
	return errors.New(errorMessage)
}
