package main

import (
	"strings"
)

type guess []struct {
	letter rune

	// value represents the color that should be displayed for this letter.
	// 0 = grey
	// 1 = yellow
	// 2 = green
	value int
}

// newGuess creates a guess from an input string and answer. The input string is
// assumed to be the correct length.
//
// When there are multiple uses of a letter in the guess, the status of those
// depends on the number and locations of that letter in the answer. Letters in
// the correct location (green) take precedence over letters in the wrong
// location (yellow). Letters in the wrong location are taken in order from left
// to right. Some examples to illustrate:
//
//     answer=aloft input=boots
//
// The first O in boots should be grey and the second should be green.
//
//     answer=aloft input=balls
//
// The first L in balls should be yellow and the second should be grey.
//
//     answer=sissy input=asses
//
// The first S is yellow, second is green, and third is yellow.
func newGuess(input, answer string) guess {
	g := make(guess, len(input))

	// Count letters in the answer so we know how many yellows we can put in the
	// output.
	answerMap := make(map[rune]int)
	for _, c := range answer {
		answerMap[c] += 1
	}

	// Count how many greens we've used in the output as they take priority over
	// yellows.
	for i, c := range input {
		// Add runes to populate the guess.
		g[i].letter = c
		// Mark all green letters so we can correctly mark yellows later.
		if c == []rune(answer)[i] {
			g[i].value = 2
			answerMap[c] -= 1
		}
	}

	for i, c := range input {
		if g[i].value == 2 {
			// Skip letters already marked as green.
			continue
		}
		if answerMap[c] > 0 {
			// If there are more of the current letter in the answer than we
			// have green letters in the current guess mark as yellow.
			g[i].value = 1
			answerMap[c] -= 1
		}
	}

	return g
}

// string builds the guess into a comparable string.
func (g guess) String() string {
	var b strings.Builder

	for _, v := range g {
		b.WriteRune(v.letter)
	}

	return b.String()
}
