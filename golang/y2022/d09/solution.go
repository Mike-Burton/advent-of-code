package d09

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Mike-Burton/advent-of-code/hashset"

	"github.com/Mike-Burton/advent-of-code/helpers"
)

// all cords are in (y, x) format

type move struct {
	dir   rune
	steps int
}

type coordinates struct {
	y, x int
}

var (
	dirToDelta = map[rune]coordinates{
		'U': {1, 0},
		'R': {0, 1},
		'D': {-1, 0},
		'L': {0, -1},
	}
	nextPos = map[coordinates]coordinates{
		// touching, do not move
		{-1, -1}: {0, 0},
		{-1, 0}:  {0, 0},
		{-1, 1}:  {0, 0},
		{0, -1}:  {0, 0},
		{0, 0}:   {0, 0},
		{0, 1}:   {0, 0},
		{1, -1}:  {0, 0},
		{1, 0}:   {0, 0},
		{1, 1}:   {0, 0},
		// move  U
		{2, 0}: {1, 0},
		// move UR
		{2, 1}: {1, 1},
		{2, 2}: {1, 1},
		{1, 2}: {1, 1},
		// move  R
		{0, 2}: {0, 1},
		// move RD
		{-1, 2}: {-1, 1},
		{-2, 2}: {-1, 1},
		{-2, 1}: {-1, 1},
		// move  D
		{-2, 0}: {-1, 0},
		// move DL
		{-2, -1}: {-1, -1},
		{-2, -2}: {-1, -1},
		{-1, -2}: {-1, -1},
		// move  L
		{0, -2}: {0, -1},
		// move LU
		{1, -2}: {1, -1},
		{2, -2}: {1, -1},
		{2, -1}: {1, -1},
	}
)

// PartOne solves the first problem of day 9 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {
	moves, err := movesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	head := coordinates{0, 0}
	tail := coordinates{0, 0}
	taiLLocs := hashset.New(tail)

	for _, m := range moves {
		headDelta := dirToDelta[m.dir]

		// for each step
		for i := 0; i < m.steps; i++ {
			// update head loc
			head = coordinates{head.y + headDelta.y, head.x + headDelta.x}

			// update tail loc
			tailDelta := nextPos[coordinates{head.y - tail.y, head.x - tail.x}]
			tail = coordinates{tail.y + tailDelta.y, tail.x + tailDelta.x}

			// track tail locs
			taiLLocs.Add(tail)
		}
	}

	_, err = fmt.Fprintf(w, "%d", taiLLocs.Size())
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day 9 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	moves, err := movesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	head := coordinates{0, 0}
	knots := [9]coordinates{}
	taiLLocs := hashset.New(knots[8])

	for _, m := range moves {
		headDelta := dirToDelta[m.dir]

		// for each step
		for i := 0; i < m.steps; i++ {
			// update head loc
			head = coordinates{head.y + headDelta.y, head.x + headDelta.x}

			for j := 0; j < len(knots); j++ {
				var knotDelta coordinates
				if j == 0 {
					knotDelta = nextPos[coordinates{head.y - knots[j].y, head.x - knots[j].x}]
				} else {
					knotDelta = nextPos[coordinates{knots[j-1].y - knots[j].y, knots[j-1].x - knots[j].x}]
				}

				// exit if no movement
				if knotDelta.y == 0 && knotDelta.x == 0 {
					break
				}

				// update knot loc
				knots[j] = coordinates{knots[j].y + knotDelta.y, knots[j].x + knotDelta.x}
			}

			// track tail locs
			taiLLocs.Add(knots[8])
		}
	}

	_, err = fmt.Fprintf(w, "%d", taiLLocs.Size())
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func movesFromReader(r io.Reader) (moves []move, err error) {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	for _, l := range lines {
		s := strings.Split(l, " ")
		steps, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, fmt.Errorf("could convert to int: %s", s[1])
		}
		moves = append(moves, move{rune(s[0][0]), steps})
	}

	return
}
