package main

import (
	"io"
	"strings"

	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/matrix"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

func ParseShape(shape string) (int, *matrix.Matrix[byte]) {
	parts := strings.SplitN(shape, ":", 2)
	id := utils.MustAtoi(parts[0])
	m := utils.Must(matrix.NewFromReader(strings.NewReader(strings.TrimSpace(parts[1]))))
	return id, m
}

func ParseRegion(region string) (matrix.Vec, []int) {
	region = strings.Replace(region, ":", "", 1)
	parts := strings.Split(region, " ")
	dims := strings.SplitN(parts[0], "x", 2)
	vec := matrix.Vec{
		X: utils.MustAtoi(dims[0]),
		Y: utils.MustAtoi(dims[1]),
	}
	presents := make([]int, 0, len(parts))
	for _, p := range parts[1:] {
		presents = append(presents, utils.MustAtoi(p))
	}
	return vec, presents
}

func PresentArea(m *matrix.Matrix[byte]) int {
	return iterators.Reduce[int, matrix.Coord](0, m.Coords(), func(acc int, item matrix.Coord) int {
		if m.AtCoord(item) == '#' {
			acc++
		}
		return acc
	})
}

func TryToFit(dims matrix.Vec, presents []int, shapes map[int]*matrix.Matrix[byte]) bool {
	area := dims.X * dims.Y
	stage.Println("area", area)
	pAreas := 0
	for i, qty := range presents {
		pArea := PresentArea(shapes[i])
		stage.Println(qty, "presents", i, "of area", pArea)
		pAreas += pArea * qty
	}
	stage.Println("pareas", pAreas, "<", area, "=>", pAreas < area)
	return pAreas < area
}

func ParseInput(input io.Reader) int {
	data := string(utils.Must(io.ReadAll(input)))
	parts := strings.Split(data, "\n\n")
	shapes := make(map[int]*matrix.Matrix[byte])
	for _, p := range parts[:len(parts)-1] {
		id, shape := ParseShape(p)
		stage.Println("Shape:", id)
		stage.Println(matrix.SMatrix(shape))
		shapes[id] = shape
	}
	res := 0
	for p := range strings.SplitSeq(parts[len(parts)-1], "\n") {
		if len(p) == 0 {
			continue
		}
		dim, presents := ParseRegion(p)
		stage.Println(dim, presents)
		fit := TryToFit(dim, presents, shapes)
		if fit {
			res++
		}
	}
	return res
}

func Stage1(input io.Reader) (any, error) {
	res := ParseInput(input)
	return res, nil
}

func Stage2(input io.Reader) (any, error) {
	return nil, nil
}
