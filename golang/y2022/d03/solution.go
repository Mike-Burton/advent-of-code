package d03

import (
	"fmt"
	"io"

	"github.com/Mike-Burton/advent-of-code/helpers"
)

type rucksack struct {
	h map[byte]struct{}
}

func (self *rucksack) Len() int {
	return len(self.h)
}

func (self *rucksack) Intersection(set *rucksack) *rucksack {
	if self.Len() < set.Len() {
		return intersection(self, set)
	}
	return intersection(set, self)
}

func intersection(a, b *rucksack) *rucksack {
	i := make(map[byte]struct{})

	for k := range a.h {
		if _, okay := b.h[k]; okay {
			i[k] = struct{}{}
		}
	}
	return &rucksack{i}
}

type splitRucksack struct {
	l rucksack
	r rucksack
}

// PartOne solves the first problem of day 3 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {
	sacks, err := splitRucksacksFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	sum := 0
	for _, v := range sacks {
		p := v.r.Intersection(&v.l)
		for k := range p.h {
			sum += itemToPriority(k)
			break
		}
	}

	_, err = fmt.Fprintf(w, "%d", sum)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day 3 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	rss, err := rucksacksFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	// double intersection
	i := 2
	sum := 0
	for i < len(rss) {
		m := rss[i-2].Intersection(&rss[i-1])
		m = m.Intersection(&rss[i])
		for k := range m.h {
			sum += itemToPriority(k)
		}
		i += 3
	}

	_, err = fmt.Fprintf(w, "%d", sum)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func splitRucksacksFromReader(r io.Reader) (rss []splitRucksack, err error) {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	//var rs []splitRucksack
	for _, l := range lines {
		rs := splitRucksack{
			r: rucksack{make(map[byte]struct{})},
			l: rucksack{make(map[byte]struct{})},
		}

		half := len(l) / 2
		for i, c := range l {
			b := byte(c)

			if i < half {
				rs.r.h[b] = struct{}{}
				continue
			}
			rs.l.h[b] = struct{}{}
		}

		rss = append(rss, rs)
	}

	return
}

func rucksacksFromReader(r io.Reader) (rss []rucksack, err error) {
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	//var rs []splitRucksack
	for _, l := range lines {
		rs := rucksack{make(map[byte]struct{})}

		half := len(l) / 2
		for i, c := range l {
			b := byte(c)

			if i < half {
				rs.h[b] = struct{}{}
				continue
			}
			rs.h[b] = struct{}{}
		}

		rss = append(rss, rs)
	}

	return
}

func itemToPriority(b byte) int {
	if byte('a') <= b && b <= byte('z') {
		return int(b) - 96
	}
	return int(b) - 64 + 26
}
