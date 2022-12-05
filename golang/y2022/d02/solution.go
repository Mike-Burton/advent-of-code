package d02

import (
	"fmt"
	"io"

	"github.com/Mike-Burton/advent-of-code/helpers"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
	LOOSE    = 0
	DRAW     = 3
	WIN      = 6
)

var scoreMap = map[string]int{
	"A X": ROCK + DRAW,
	"A Y": PAPER + WIN,
	"A Z": SCISSORS + LOOSE,
	"B X": ROCK + LOOSE,
	"B Y": PAPER + DRAW,
	"B Z": SCISSORS + WIN,
	"C X": ROCK + WIN,
	"C Y": PAPER + LOOSE,
	"C Z": SCISSORS + DRAW,
}

var riggedMap = map[string]int{
	"A X": LOOSE + SCISSORS,
	"A Y": DRAW + ROCK,
	"A Z": WIN + PAPER,
	"B X": LOOSE + ROCK,
	"B Y": DRAW + PAPER,
	"B Z": WIN + SCISSORS,
	"C X": LOOSE + PAPER,
	"C Y": DRAW + SCISSORS,
	"C Z": WIN + ROCK,
}

// PartOne solves the first problem of day 2 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	score := 0
	for _, l := range lines {
		score += scoreMap[l]
	}

	_, err = fmt.Fprintf(w, "%d", score)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day 2 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	score := 0
	for _, l := range lines {
		score += riggedMap[l]
	}

	_, err = fmt.Fprintf(w, "%d", score)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}
