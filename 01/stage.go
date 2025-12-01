package main

import (
	"io"

	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

func Stage1(input io.Reader) (any, error) {
	idx := 50
	zeros := 0
	for line := range iterators.MustLines(input) {
		stage.Println(line)
		mul := 0
		switch line[0] {
		case 'L':
			mul = -1
		case 'R':
			mul = 1
		default:
			panic(line[0])
		}
		val := utils.MustAtoi(line[1:]) * mul
		stage.Println("val", val)
		idx += val
		idx = (idx + 100) % 100
		stage.Println("idx", idx)
		if idx == 0 {
			zeros++
		}
	}
	return zeros, nil
}

func Stage2(input io.Reader) (any, error) {
	idx := 50
	zeros := 0
	for line := range iterators.MustLines(input) {
		stage.Println()
		stage.Println(line)
		mul := 0
		switch line[0] {
		case 'L':
			mul = -1
		case 'R':
			mul = 1
		default:
			panic(line[0])
		}
		diff := utils.MustAtoi(line[1:]) * mul
		// Checks
		if diff == 0 {
			panic("zero case is not supported")
		}
		// Checks
		oldidx := idx
		idx = (((idx + diff) % 100) + 100) % 100
		stage.Println("move", oldidx, "--(", diff, ")->", idx)
		// added := false
		zerosToAdd := (diff * mul) / 100
		if zerosToAdd > 0 {
			stage.Println("add", zerosToAdd, "for full turns")
		}
		if mul == 1 {
			// Turning Right
			if idx < oldidx {
				stage.Println("add", 1, "for crossing")
				zerosToAdd++
			}
		} else {
			// Turning Left
			if oldidx != 0 && idx > oldidx {
				stage.Println("add", 1, "for crossing")
				zerosToAdd++
			} else if oldidx != 0 && idx == 0 {
				stage.Println("add", 1, "for exact stop")
				zerosToAdd++
			}
		}
		zeros += zerosToAdd
		stage.Println("add", zerosToAdd, "total")
		// if !added && idx == 0 {
		// 	zeros++
		// }
	}
	return zeros, nil
}
