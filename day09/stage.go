package main

import (
	"io"
	"math"
	"strings"

	"github.com/nlm/adventofcode2025/internal/combinations"
	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/matrix"
	"github.com/nlm/adventofcode2025/internal/sets"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

func Dist(a, b matrix.Coord) float64 {
	return math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2))
}

func Area(a, b matrix.Coord) int {
	area := 1
	if a.X > b.X {
		area *= a.X - b.X + 1
	} else {
		area *= b.X - a.X + 1
	}
	if a.Y > b.Y {
		area *= a.Y - b.Y + 1
	} else {
		area *= b.Y - a.Y + 1
	}
	return area
}

func FindMaxArea(coords []matrix.Coord) int {
	pairs := func(yield func([2]matrix.Coord) bool) {
		for pair := range combinations.Combinations(coords, 2) {
			if !yield([2]matrix.Coord{pair[0], pair[1]}) {
				return
			}
		}
	}
	return iterators.Reduce(0, pairs, func(m int, coords [2]matrix.Coord) int {
		area := Area(coords[0], coords[1])
		if area > m {
			return area
		}
		return m
	})
}

// A..C
// ....
// D..B
func FindOtherCorners(a, b matrix.Coord) (matrix.Coord, matrix.Coord) {
	c := matrix.Coord{X: a.X, Y: b.Y}
	d := matrix.Coord{X: b.X, Y: a.Y}
	return c, d
}

func FindMaxArea2(coords []matrix.Coord) int {
	// make a set of Coords
	coordsSet := make(sets.Set[matrix.Coord], len(coords))
	sets.Append(coordsSet, coords...)

	pairs := func(yield func([2]matrix.Coord) bool) {
		for pair := range combinations.Combinations(coords, 2) {
			if !yield([2]matrix.Coord{pair[0], pair[1]}) {
				return
			}
		}
	}

	return iterators.Reduce(0, pairs, func(m int, pair [2]matrix.Coord) int {
		stage.Println()
		a, b := pair[0], pair[1]

		// calculate area
		area := Area(a, b)
		stage.Println("current max:", m)
		stage.Println("area", a, b, "->", area)
		// area not interesting
		if area < m {
			return m
		}
		// find intersecting lines
		p1 := matrix.Coord{X: min(a.X, b.X), Y: min(a.Y, b.Y)}
		p2 := matrix.Coord{X: max(a.X, b.X), Y: max(a.Y, b.Y)}
		for i := range len(coords) {
			c1 := coords[i]
			c2 := coords[(i+1)%len(coords)]
			if !(max(c1.X, c2.X) <= p1.X ||
				p2.X <= min(c1.X, c2.X) ||
				max(c1.Y, c2.Y) <= p1.Y ||
				p2.Y <= min(c1.Y, c2.Y)) {
				return m
			}
		}
		return area
	})
}

func Stage1(input io.Reader) (any, error) {
	coords := make([]matrix.Coord, 0)
	for line := range iterators.MustLines(input) {
		parts := strings.Split(line, ",")
		coords = append(coords, matrix.Coord{X: utils.MustAtoi(parts[0]), Y: utils.MustAtoi(parts[1])})
	}
	stage.Println(coords)
	maxArea := FindMaxArea(coords)
	return maxArea, nil
}

// greater than 140762692
func Stage2(input io.Reader) (any, error) {
	coords := make([]matrix.Coord, 0)
	for line := range iterators.MustLines(input) {
		parts := strings.Split(line, ",")
		coords = append(coords, matrix.Coord{X: utils.MustAtoi(parts[0]), Y: utils.MustAtoi(parts[1])})
	}
	stage.Println(coords)
	maxArea := FindMaxArea2(coords)
	return maxArea, nil
}
