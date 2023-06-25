package words_test

import (
	"testing"

	"github.com/austien/type-of-the-bored/ansi"
	"github.com/austien/type-of-the-bored/words"
)

func TestToString(t *testing.T) {
	letterA := words.NewLetter('a')

	letterB := words.NewLetter('b')
	letterB.Color = ansi.RED
	letterB.HasUnderline = true

	letterC := words.NewLetter('c')
	letterC.Color = ansi.GREEN

	// Run the test with the verbose flag (-v) to see this output
	t.Logf("%s%s%s\n", letterA.ToString(), letterB.ToString(), letterC.ToString())

}
