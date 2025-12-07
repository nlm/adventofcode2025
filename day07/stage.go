package main

import (
	"io"

	"github.com/nlm/adventofcode2025/internal/matrix"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

func MatrixSearch[T comparable](m *matrix.Matrix[T], val T) (matrix.Coord, bool) {
	for c := range m.Coords() {
		if m.AtCoord(c) == val {
			return c, true
		}
	}
	return matrix.Coord{}, false
}

func Stage1(input io.Reader) (any, error) {
	m := utils.Must(matrix.NewFromReader(input))
	split := 0
	for c := range m.Coords() {
		switch m.AtCoord(c) {
		case '.':
			up := c.Up()
			if m.InCoord(up) && (m.AtCoord(up) == '|' || m.AtCoord(up) == 'S') {
				m.SetAtCoord(c, '|')
			}
		case '^':
			up := c.Up()
			if m.InCoord(up) && m.AtCoord(up) == '|' {
				split++
				for _, c2 := range []matrix.Coord{c.Left(), c.Right()} {
					if m.InCoord(c2) && m.AtCoord(c2) == '.' {
						m.SetAtCoord(c2, '|')
					}
				}
			}
		}
	}
	stage.Println(matrix.SMatrix(m))
	return split, nil
}

func Stage2(input io.Reader) (any, error) {
	m := utils.Must(matrix.NewFromReader(input))
	stage.Println(matrix.SMatrix(m))
	nm := matrix.New[int](m.Size.X, m.Size.Y)
	for c := range m.Coords() {
		switch m.AtCoord(c) {
		case 'S':
			nm.SetAtCoord(c, 1)
		case '.':
			up := c.Up()
			if m.InCoord(up) && (m.AtCoord(up) == '|' || m.AtCoord(up) == 'S') {
				m.SetAtCoord(c, '|')
				nm.SetAtCoord(c, nm.AtCoord(up))
			}
		case '^':
			up := c.Up()
			if m.InCoord(up) && m.AtCoord(up) == '|' {
				nm.SetAtCoord(c, nm.AtCoord(up))
				for _, c2 := range []matrix.Coord{c.Left(), c.Right()} {
					if m.InCoord(c2) && (m.AtCoord(c2) == '.' || m.AtCoord(c2) == '|') {
						m.SetAtCoord(c2, '|')
						nm.SetAtCoord(c2, nm.AtCoord(c)+nm.AtCoord(c2))
						if c2 == c.Right() {
							// add upper value when looking right, as the previous case ('.') did not trigger yet
							nm.SetAtCoord(c2, nm.AtCoord(c2)+nm.AtCoord(c2.Up()))
						}
					}
				}
			}
		}
	}
	stage.Print(matrix.SMatrix(m))
	for y := range nm.Size.Y {
		for x := range nm.Size.X {
			v := nm.At(x, y)
			if m.At(x, y) == '^' {
				stage.Print("^")
			} else if v > 0 {
				stage.Print(v)
			} else {
				stage.Print(".")
			}
		}
		stage.Print("\n")
	}
	total := 0
	for x := range nm.Size.X {
		value := nm.At(x, nm.Size.Y-1)
		total += value
	}
	return total, nil
}
