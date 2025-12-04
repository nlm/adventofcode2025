package main

import (
	"io"
	"iter"

	"github.com/nlm/adventofcode2025/internal/matrix"
	"github.com/nlm/adventofcode2025/internal/utils"
)

func Around(m *matrix.Matrix[byte], coord matrix.Coord) iter.Seq[matrix.Coord] {
	return func(yield func(matrix.Coord) bool) {
		for _, vec := range []matrix.Vec{
			matrix.UpLeft,
			matrix.Up,
			matrix.UpRight,
			matrix.Right,
			matrix.DownRight,
			matrix.Down,
			matrix.DownLeft,
			matrix.Left,
		} {
			nc := coord.Add(vec)
			// if !m.InCoord(nc) {
			// 	continue
			// }
			if !yield(nc) {
				return
			}
		}
	}
}

func Stage1(input io.Reader) (any, error) {
	m := utils.Must(matrix.NewFromReader(input))
	accessible := 0
	for c := range m.Coords() {
		if m.AtCoord(c) != '@' {
			// not a roll of paper
			continue
		}
		adjRolls := 0
		for ac := range Around(m, c) {
			if m.InCoord(ac) && m.AtCoord(ac) == '@' {
				adjRolls++
			}
		}
		if adjRolls < 4 {
			accessible++
		}
	}
	return accessible, nil
}

func Stage2(input io.Reader) (any, error) {
	m := utils.Must(matrix.NewFromReader(input))
	total := 0
	for updated := true; updated; {
		updated = false
		for c := range m.Coords() {
			if m.AtCoord(c) != '@' {
				// not a roll of paper
				continue
			}
			adjRolls := 0
			for ac := range Around(m, c) {
				if m.InCoord(ac) && m.AtCoord(ac) == '@' {
					adjRolls++
				}
			}
			if adjRolls < 4 {
				// can be removed
				m.SetAtCoord(c, '.')
				total++
				// new items could be available
				updated = true
			}
		}
	}
	return total, nil
}
