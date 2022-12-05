package d05

import (
	"fmt"
	"github.com/busser/adventofcode/helpers"
	"github.com/busser/adventofcode/stack"
	"io"
	"strconv"
	"strings"
)

type rearrangement struct {
	move int
	from int
	to   int
}

// PartOne solves the first problem of day 5 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {
	cargo, moves, err := cargoFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	// make moves
	for _, m := range moves {
		for i := 0; i < m.move; i++ {
			crate := cargo[m.from].Pop()
			cargo[m.to].Push(crate)
		}
	}

	_, err = fmt.Fprintf(w, "%s", cargoTops(cargo))
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day 5 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	cargo, moves, err := cargoFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	// make moves
	for _, m := range moves {
		var crates []rune

		// grab N crates
		for i := 0; i < m.move; i++ {
			crates = append(crates, cargo[m.from].Pop())
		}

		// place N crates
		for i := m.move - 1; i > -1; i-- {
			cargo[m.to].Push(crates[i])
		}
	}

	_, err = fmt.Fprintf(w, "%s", cargoTops(cargo))
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func cargoFromReader(r io.Reader) (cargo []stack.Stack[rune], moves []rearrangement, err error) {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read input: %w", err)
	}

	// init cargo stacks
	n := len(lines[0])/4 + 1
	for i := 0; i < n; i++ {
		cargo = append(cargo, *stack.New[rune]())
	}

	// find breakpoint between cargo and move
	breakPoint := 0
	for ; ; breakPoint++ {
		if len(lines[breakPoint]) == 0 {
			break
		}
	}

	// read in cargo
	for i := breakPoint - 2; i != -1; i-- {
		stackPos := 1
		for j := 0; j < n; j++ {
			s := rune(lines[i][stackPos])
			if s != ' ' {
				cargo[j].Push(s)
			}

			stackPos += 4
		}
	}

	// read in moves
	for i := breakPoint + 1; i < len(lines); i++ {
		var (
			ss   = strings.Split(lines[i], " ")
			m, _ = strconv.Atoi(ss[1])
			f, _ = strconv.Atoi(ss[3])
			t, _ = strconv.Atoi(ss[5])
		)

		// adjust from and to for zero index
		moves = append(moves, rearrangement{m, f - 1, t - 1})
	}

	return
}

func cargoTops(cargo []stack.Stack[rune]) string {
	var ret []rune
	for _, s := range cargo {
		if s.Size() == 0 {
			continue
		}
		ret = append(ret, s.Peek())
	}
	return string(ret)
}
