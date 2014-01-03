package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Adjectives
var adjectives = []string{
	// Leaked code names
	"main",    // MAINSTAGE
	"flying",  // FLYING PIG
	"quantum", // QUANTUM INSERT
	"ferret",  // FERRET CANNON
	//"love",  // LOVEINT

	// Generic colors
	"blue",
	"green",
	//"orange",
	"red",
	//"violet",
	"yellow",

	// From NSAGENERATOR
	/*
	"loud",
	"irate",
	"angry",
	"peeved",
	"happy",
	"slimy",
	"sleepy",
	"junior",
	"slicker",
	"united",
	"somber",
	"bizarre",
	"odd",
	"weird",
	"wrong",
	"latent",
	"chilly",
	"strange",
	"loud",
	"silent",
	"hopping",
	"violent",
	*/
}

var nouns = []string{
	"pig",    // FLYINGPIG
	"insert", // QUANTUM INSERT
	"cannon", // FERRET CANNON
	//"int",  // LOVEINT

	// From NSAGENERATOR
	"whisper",
	//"felony",
	"moon",
	"sucker",
	"penguin",
	"waffle",
	//"maestro",
	"night",
	"trinity",
	//"deity",
	"monkey",
	"ark",
	"squirrel",
	"iron",
	"bounce",
	"farm",
	"chef",
	"trough",
	"net",
	"trawl",
	"glee",
	"water",
	"spork",
	"plow",
	"feed",
	"souffle",
	"route",
	"bagel",
	"montana",
	"analyst",
	"auto",
	"watch",
	"photo",
	"yard",
	"source",
	"monkey",
	"seagull",
	"toll",
	"spawn",
	"gopher",
	//"chipmunk",
	"set",
	"calendar",
	"artist",
	"chaser",
	"scan",
	"tote",
	"beam",
	//"entourage",
	//"genesis",
	"walk",
	"spatula",
	//"rage",
	"fire",
	"master",
}

func randomCodeword() string {
	return adjectives[rand.Intn(len(adjectives))] + nouns[rand.Intn(len(nouns))]
}

func main() {
	rand.Seed(time.Now().Unix())
	// Print in three columns
	words := make([]string, 30)
	longest := 0
	for i := 0; i<30; i++ {
		words[i] = randomCodeword()
		if len(words[i]) > longest {
			longest = len(words[i])
		}
	}

	fmtString := fmt.Sprintf("%%-%ds %%-%ds %%s\n", longest + 1, longest + 1);
	for i := 0; i<30; i += 3 {
		fmt.Printf(fmtString, words[i], words[i + 1], words[i + 2]);
	}
}
