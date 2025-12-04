package main

import (
	"io"
	"slices"
	"strings"

	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/stage"
)

func StrToInts(s string) []int {
	// return slices.Collect(iterators.Map(slices.Values([]byte(s)), func(b byte) int {
	// 	return int(b - '0')
	// }))
	ints := make([]int, 0, len(s))
	for _, c := range s {
		ints = append(ints, int(c-'0'))
	}
	return ints
}

func FindMax(ints []int) int {
	return iterators.Reduce(0, slices.Values(ints), func(a, b int) int {
		if a > b {
			return a
		}
		return b
	})
}

func HandleLine1(line string) int {
	ints := StrToInts(line)
	stage.Println(line, "->", ints)
	maxDozen := FindMax(ints[:len(ints)-1])
	maxIdx := strings.IndexRune(line, '0'+rune(maxDozen))
	maxUnit := FindMax(ints[maxIdx+1:])
	stage.Println("max", maxDozen, "maxUnit", maxUnit)
	return maxDozen*10 + maxUnit
}

func Stage1(input io.Reader) (any, error) {
	res := 0
	for line := range iterators.MustLines(input) {
		value := HandleLine1(line)
		stage.Println(value)
		res += value
	}
	return res, nil
}

func FindMaxIdx(ints []int) (maxVal int, idx int) {
	maxVal = FindMax(ints)
	return maxVal, slices.Index(ints, maxVal)
}

func HandleLine2(line string, intLen int) int {
	ints := StrToInts(line)
	stage.Println(line, "->", ints)
	leftIdx := 0
	total := 0
	for i := intLen - 1; i >= 0; i-- {
		val, idx := FindMaxIdx(ints[leftIdx : len(ints)-i])
		leftIdx += idx + 1
		total = total*10 + val
	}
	return total
}

func Stage2(input io.Reader) (any, error) {
	res := 0
	for line := range iterators.MustLines(input) {
		value := HandleLine2(line, 12)
		stage.Println(value)
		res += value
	}
	return res, nil
}
