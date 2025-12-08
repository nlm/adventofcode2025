package main

import (
	"fmt"
	"io"
	"maps"
	"math"
	"slices"
	"strings"

	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/sets"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

type Box [3]int

func (b Box) X() int {
	return b[0]
}

func (b Box) Y() int {
	return b[1]
}

func (b Box) Z() int {
	return b[2]
}

func (b Box) Dist(b2 Box) float64 {
	return math.Sqrt(math.Pow(float64(b.X()-b2.X()), 2) + math.Pow(float64(b.Y()-b2.Y()), 2) + math.Pow(float64(b.Z()-b2.Z()), 2))
}

type Circuits []sets.Set[Box]

func (cs Circuits) Count() int {
	return len(cs)
}

func (cs *Circuits) Connect(b1, b2 Box) {
	var b1Set, b2Set sets.Set[Box]
	var b1Idx, b2Idx int
	for i, c := range *cs {
		if b1Set == nil && c.Contains(b1) {
			b1Set = c
			b1Idx = i
		} else if b2Set == nil && c.Contains(b2) {
			b2Set = c
			b2Idx = i
		}
		if b1Set != nil && b2Set != nil {
			break
		}
	}
	_, _ = b1Idx, b2Idx
	if b1Set == nil && b2Set == nil {
		stage.Println("connect: all nil, new set")
		newSet := make(sets.Set[Box])
		newSet.Add(b1, b2)
		*cs = append(*cs, newSet)
	} else if b1Set != nil && b2Set != nil {
		stage.Println("connect: no nil, merge sets")
		sets.Insert(b1Set, sets.Values(b2Set))
		*cs = slices.Delete(*cs, b2Idx, b2Idx+1)
		clear(b2Set)
	} else if b1Set != nil {
		stage.Println("connect: add", b2, "to existing", b1, "set")
		b1Set.Add(b2)
	} else if b2Set != nil {
		stage.Println("connect: add", b1, "to existing", b2, "set")
		b2Set.Add(b1)
	} else {
		panic("internal error")
	}
}

func CalculateDistances(boxes []Box) map[float64]*[2]Box {
	dists := make(map[float64]*[2]Box)
	for _, b1 := range boxes {
		for _, b2 := range boxes {
			if b1 == b2 {
				continue
			}
			dist := b1.Dist(b2)
			if dists[dist] != nil {
				if *dists[dist] != [2]Box{b2, b1} {
					panic("duplicate")
				}
				continue
			}
			dists[dist] = &[2]Box{b1, b2}
		}
	}
	return dists
}

func ParseBoxes(input io.Reader) []Box {
	boxes := make([]Box, 0)
	for line := range iterators.MustLines(input) {
		parts := strings.Split(line, ",")
		boxes = append(boxes, Box{utils.MustAtoi(parts[0]), utils.MustAtoi(parts[1]), utils.MustAtoi(parts[2])})
	}
	return boxes
}

func Stage1(input io.Reader) (any, error) {
	boxes := ParseBoxes(input)
	// stage.Println(boxes)

	// Warning: auto-adjust for examples
	var maxIterations = 1000
	if len(boxes) < 1000 {
		maxIterations = 10
	}
	dists := CalculateDistances(boxes)
	// dists := map[float64]*[2]Box{
	// 	1: {{1, 1, 1}, {2, 2, 2}},
	// 	2: {{3, 3, 3}, {4, 4, 4}},
	// 	3: {{5, 5, 5}, {4, 4, 4}},
	// }
	circuits := make(Circuits, 0)
	iterations := 0
	for _, k := range slices.Sorted(maps.Keys(dists)) {
		stage.Println("handling:", k, dists[k])
		circuits.Connect((*dists[k])[0], (*dists[k])[1])
		stage.Println("circuits after:", circuits)
		iterations++
		if !(iterations < maxIterations) {
			break
		}
		stage.Println()
	}
	circuitLens := iterators.MapSlice(circuits, func(s sets.Set[Box]) int {
		return len(s)
	})
	slices.SortFunc(circuitLens, func(a, b int) int {
		return b - a
	})
	stage.Println("lengths:", circuitLens)
	result := 1
	for i, v := range circuitLens {
		if !(i < 3) {
			break
		}
		result *= v
	}
	return result, nil
}

func Stage2(input io.Reader) (any, error) {
	boxes := ParseBoxes(input)
	dists := CalculateDistances(boxes)
	circuits := make(Circuits, 0)
	for _, k := range slices.Sorted(maps.Keys(dists)) {
		stage.Println("handling:", "dist=", k, "set=", dists[k])
		circuits.Connect((*dists[k])[0], (*dists[k])[1])
		// stage.Println("circuits after:", circuits)
		stage.Println("CHECK:", len(boxes), len(circuits[0]))
		if len(boxes) == len(circuits[0]) {
			b1, b2 := (*dists[k])[0], (*dists[k])[1]
			stage.Println("Final CNX:", b1, b2, "->", b1.X()*b2.X())
			return b1.X() * b2.X(), nil
		}
		stage.Println()
	}
	return nil, fmt.Errorf("no solution found")
}
