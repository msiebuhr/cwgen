package main

import (
	"fmt"
	"math/rand"
	"time"
)

type wordList struct {
	words []string
	occurrences []uint
}

func NewWordList() *wordList {
	return &wordList{
		words: make([]string, 0),
		occurrences: make([]uint, 0),
	}
}

func (w *wordList) addWord(newWord string) {
	// Loop through the words
	var index int = -1;
	for i, word := range w.words {
		if word == newWord {
			index = i;
			break
		}
	}

	if index == -1 {
		index = len(w.words);
		w.words = append(w.words, newWord)
		w.occurrences = append(w.occurrences, 0);
	}

	w.occurrences[index] += 1;
}

func (w *wordList) addWords(words []string) {
	for _, word := range words {
		w.addWord(word)
	}
}

func (w *wordList) getRandomWord() string {
	return w.words[rand.Intn(len(w.words))]
}

func (w *wordList) getWeighedRandomWord() string {
	var sum uint = 0;
	for _, occ := range w.occurrences {
		sum += occ
	}

	index := uint(rand.Intn(int(sum)))

	// Look what word that lands us at
	for i, occ := range w.occurrences {
		if occ > index {
			return w.words[i]
		} else {
			index -= occ
		}
	}

	return ""
}


// For whole words
type codewordGen struct {
	prefix, adjective, infix, noun, suffix *wordList
}

func NewCodewordGen() *codewordGen {
	return &codewordGen{
		prefix: NewWordList(),
		adjective: NewWordList(),
		infix: NewWordList(),
		noun: NewWordList(),
		suffix: NewWordList(),
	}
}

func (c *codewordGen) addCodeword(prefix, adj, infix, noun, suffix string) {
	c.prefix.addWord(prefix)
	c.adjective.addWord(adj)
	c.infix.addWord(infix)
	c.noun.addWord(noun)
	c.suffix.addWord(suffix)
}

// TODO: Strip output
func (c *codewordGen) getCodeword() string {
	return c.prefix.getWeighedRandomWord() + c.adjective.getRandomWord() + c.infix.getWeighedRandomWord() + c.noun.getRandomWord() + c.suffix.getWeighedRandomWord()
}

func main() {
	rand.Seed(time.Now().Unix())
	// Print in three columns
	words := make([]string, 30)
	longest := 0

	gen := NewCodewordGen()

	// Single word
	//gen.addCodeword("", "", " ", "ECHELON", "")
	//gen.addCodeword("", "", " ", "PRISM", "")

	// Multi word
	gen.addCodeword("", "AIR", "", "GAP", "")
	gen.addCodeword("", "AIR", "", "HANDLER", "")
	gen.addCodeword("", "AMBER", "", "JACK", "")
	gen.addCodeword("", "BELL", " ", "TOPPER", "")
	gen.addCodeword("", "BLIND", " ", "DATE", "")
	gen.addCodeword("", "BULL", "", "RUN", "")
	gen.addCodeword("", "COURIER", "", "SKILL", "")
	gen.addCodeword("", "CROSS", "", "BONES", "")
	gen.addCodeword("", "DEW", "", "SWEEPER", "")
	gen.addCodeword("", "DISH", "", "FIRE", "")
	gen.addCodeword("", "DROP", "", "MIRE", "")
	gen.addCodeword("", "DROPOUT", "", "JEEP", "")
	gen.addCodeword("", "DRY", "", "TORTUGAS", "")
	gen.addCodeword("", "FALLEN", "", "ORACLE", "")
	gen.addCodeword("", "FEED", "", "THROUGH", "")
	gen.addCodeword("", "FERRET", " ", "CANNON", "")
	gen.addCodeword("", "FIRE", "", "WALK", "")
	gen.addCodeword("", "FLUX", "", "RABBIT", "")
	gen.addCodeword("", "FLYING", " ", "PIG", "")
	gen.addCodeword("", "FLYING", " ", "PIG", "")
	gen.addCodeword("", "FOX", "", "ACID", "")
	gen.addCodeword("", "FOX", "", "SEARCH", "")
	gen.addCodeword("", "FOX", "", "TRAIL", "")
	gen.addCodeword("", "GHOST", "", "MACHINE", "")
	gen.addCodeword("", "GOURMET", "", "THROUGH", "")
	gen.addCodeword("", "GREY", "", "STONE", "")
	gen.addCodeword("", "HAMMER", "", "MILL", "")
	gen.addCodeword("", "HAPPY", "", "FOOT", "")
	gen.addCodeword("", "HAVE", "", "QUICK", "")
	gen.addCodeword("", "IRON", "", "CHEF", "")
	gen.addCodeword("", "IRON", "", "SAND", "")
	gen.addCodeword("", "MAIN", "", "STAGE", "")
	gen.addCodeword("", "QUANTUM", " ", "INSERT", "")
	gen.addCodeword("X", "KEY", "", "SCORE", "")

	for i := 0; i<30; i++ {
		words[i] = gen.getCodeword()
		if len(words[i]) > longest {
			longest = len(words[i])
		}
	}

	// Use tabwriter
	fmtString := fmt.Sprintf("%%-%ds %%-%ds %%s\n", longest + 1, longest + 1);
	for i := 0; i<30; i += 3 {
		fmt.Printf(fmtString, words[i], words[i + 1], words[i + 2]);
	}
}
