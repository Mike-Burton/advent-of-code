package d01

import (
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/busser/adventofcode/helpers"
)

// PartOne solves the first problem of day 1 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {
	lines, err := elfCalories(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	_, err = fmt.Fprintf(w, "%d", lines[0])
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day 1 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	lines, err := elfCalories(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	_, err = fmt.Fprintf(w, "%d", lines[0]+lines[1]+lines[2])
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func elfCalories(r io.Reader) ([]int, error) {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return nil, err
	}

	var elfCals []int
	thisElf := 0

	for _, l := range lines {
		if len(l) == 0 {
			elfCals = append(elfCals, thisElf)
			thisElf = 0
			continue
		}

		c, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		thisElf += c
	}

	//
	sort.Sort(sort.Reverse(sort.IntSlice(elfCals)))

	return elfCals, nil
}
