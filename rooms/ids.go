package rooms

import (
	"fmt"
	"math/rand"
)

func NewID() string {
	return fmt.Sprintf("%s-%s", Adjectives[rand.Intn(len(Adjectives))], Nouns[rand.Intn(len(Nouns))])
}

var Adjectives = []string{
	"amazing",
	"beautiful",
	"charming",
	"daring",
	"excellent",
	"fearless",
	"graceful",
	"happy",
	"incredible",
	"joyful",
	"kind",
	"lovely",
	"magnificent",
	"nice",
	"outstanding",
	"perfect",
	"quiet",
	"remarkable",
	"smart",
	"terrific",
	"unique",
	"victorious",
	"witty",
	"exciting",
	"youthful",
	"zealous",
	"brave",
	"clever",
	"elegant",
	"friendly",
	"gentle",
	"helpful",
	"intelligent",
	"jolly",
	"kind",
	"lively",
	"merry",
	"neat",
	"optimistic",
	"proud",
	"quiet",
	"resilient",
	"sweet",
	"tender",
	"upbeat",
	"vibrant",
	"warm",
	"adventurous",
	"bright",
	"cheerful",
	"determined",
	"energetic",
	"fearless",
	"gracious",
	"helpful",
	"inspiring",
	"joyful",
	"kind",
	"lucky",
	"magnetic",
	"noble",
	"optimistic",
	"patient",
	"quick-witted",
	"reliable",
	"strong",
	"talented",
	"unwavering",
	"versatile",
	"wise",
	"excellent",
	"youthful",
	"zestful",
	"amiable",
	"benevolent",
	"calm",
	"decisive",
	"eloquent",
	"faithful",
	"generous",
	"hopeful",
	"intrepid",
	"joyful",
	"keen",
	"lively",
	"mild",
	"noble",
	"optimistic",
	"practical",
	"quiet",
	"resourceful",
	"strong",
	"tactful",
	"upbeat",
	"valiant",
	"wise",
	"zealous",
}

var Nouns = []string{
	"apple",
	"banana",
	"car",
	"dog",
	"elephant",
	"flower",
	"grape",
	"hat",
	"ice",
	"jacket",
	"kiwi",
	"lemon",
	"mango",
	"nut",
	"orange",
	"pear",
	"quilt",
	"rabbit",
	"star",
	"tree",
	"umbrella",
	"vase",
	"watermelon",
	"xylophone",
	"yacht",
	"zebra",
	"airplane",
	"ball",
	"cat",
	"desk",
	"elephant",
	"fish",
	"grape",
	"horse",
	"ice",
	"jelly",
	"kite",
	"lemon",
	"mouse",
	"nest",
	"orange",
	"pen",
	"queen",
	"rose",
	"sun",
	"table",
	"umbrella",
	"violin",
	"water",
	"xylophone",
	"yarn",
	"zebra",
	"apple",
	"bear",
	"carrot",
	"dog",
	"egg",
	"flower",
	"grape",
	"hat",
	"ice",
	"juice",
	"kangaroo",
	"lemon",
	"mango",
	"nut",
	"orange",
	"pear",
	"queen",
	"rabbit",
	"star",
	"tree",
	"umbrella",
	"violin",
	"water",
	"xylophone",
	"yacht",
	"zebra",
	"airplane",
	"ball",
	"cat",
	"dog",
	"elephant",
	"fish",
	"grape",
	"horse",
	"ice",
	"jelly",
	"kite",
	"lemon",
	"mouse",
	"nest",
	"orange",
	"pear",
	"queen",
	"rose",
	"sun",
	"table",
	"umbrella",
}
