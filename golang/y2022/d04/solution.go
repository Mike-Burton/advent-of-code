package d04

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/busser/adventofcode/helpers"
)

type section struct {
	lower int
	upper int
}

type elfPair struct {
	eo section
	et section
}

// PartOne solves the first problem of day 4 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {
	elfPairs, err := elfPairsFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	sum := 0
	for _, ep := range elfPairs {
		if ep.eo.lower <= ep.et.lower && ep.eo.upper >= ep.et.upper {
			sum += 1
			continue
		}

		if ep.et.lower <= ep.eo.lower && ep.et.upper >= ep.eo.upper {
			sum += 1
			continue
		}
	}

	_, err = fmt.Fprintf(w, "%d", sum)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day 4 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	elfPairs, err := elfPairsFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	sum := 0
	for _, ep := range elfPairs {
		if ep.eo.lower >= ep.et.lower && ep.eo.lower <= ep.et.upper ||
			ep.eo.upper <= ep.et.lower && ep.eo.upper >= ep.et.upper {
			sum += 1
			continue

		}

		if ep.et.lower >= ep.eo.lower && ep.et.lower <= ep.eo.upper ||
			ep.et.upper <= ep.eo.lower && ep.et.upper >= ep.eo.upper {
			sum += 1
			continue
		}
	}

	_, err = fmt.Fprintf(w, "%d", sum)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func elfPairsFromReader(r io.Reader) (eps []elfPair, err error) {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	for _, l := range lines {
		var (
			// split on "," -> gives two elfs
			elfs = strings.Split(l, ",")
			// split on "-" -> gives lower and upper
			eor = strings.Split(elfs[0], "-")
			etr = strings.Split(elfs[1], "-")
			// parse loser and upper for each elf
			eol, _ = strconv.Atoi(eor[0])
			eou, _ = strconv.Atoi(eor[1])
			etl, _ = strconv.Atoi(etr[0])
			etu, _ = strconv.Atoi(etr[1])
		)

		// add to list
		eps = append(eps, elfPair{
			eo: section{eol, eou},
			et: section{etl, etu},
		})
	}

	return eps, nil
}
