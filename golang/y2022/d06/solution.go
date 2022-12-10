package d06

import (
	"fmt"
	"io"

	"github.com/Mike-Burton/advent-of-code/hashset"
	"github.com/Mike-Burton/advent-of-code/helpers"
)

const (
	PARTONE = 4
	PARTTWO = 14
)

// PartOne solves the first problem of day 6 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {

	chars, err := signalFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	i := PARTONE - 1
	for ; i < len(chars); i++ {
		hs := hashset.New(chars[i], chars[i-1], chars[i-2], chars[i-3])
		if hs.Size() == PARTONE {
			break
		}
	}

	_, err = fmt.Fprintf(w, "%d", i+1)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day 6 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	chars, err := signalFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	i := PARTTWO - 1
	for ; i < len(chars); i++ {
		hs := hashset.New(
			chars[i],
			chars[i-1],
			chars[i-2],
			chars[i-3],
			chars[i-4],
			chars[i-5],
			chars[i-6],
			chars[i-7],
			chars[i-8],
			chars[i-9],
			chars[i-10],
			chars[i-11],
			chars[i-12],
			chars[i-13],
		)

		if hs.Size() == PARTTWO {
			break
		}
	}
	_, err = fmt.Fprintf(w, "%d", i+1)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func signalFromReader(r io.Reader) (chars []rune, err error) {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	for _, r := range lines[0] {
		chars = append(chars, r)
	}

	return
}
