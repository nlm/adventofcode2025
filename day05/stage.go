package main

import (
	"io"
	"slices"
	"strings"

	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

type Range struct {
	Low  int
	High int
}

func (r Range) Contains(n int) bool {
	return n >= r.Low && n <= r.High
}

func (r Range) Len() int {
	return r.High - r.Low + 1
}

func ReadRangesProducts(input io.Reader) ([]Range, []int) {
	lines := slices.Collect(iterators.MustLines(input))
	i := 0
	ranges := make([]Range, 0)
	products := make([]int, 0)
	for ; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			i++
			break
		}
		parts := strings.Split(line, "-")
		ranges = append(ranges, Range{
			Low:  utils.MustAtoi(parts[0]),
			High: utils.MustAtoi(parts[1]),
		})
	}
	for ; i < len(lines); i++ {
		line := lines[i]
		products = append(products, utils.MustAtoi(line))
	}
	stage.Println("ranges", ranges)
	stage.Println("products", products)
	return ranges, products
}

func Stage1(input io.Reader) (any, error) {
	fresh := 0
	ranges, products := ReadRangesProducts(input)
	for _, product := range products {
		for _, r := range ranges {
			if r.Contains(product) {
				fresh++
				break
			}
		}
	}
	return fresh, nil
}

func Stage2(input io.Reader) (any, error) {
	ranges, _ := ReadRangesProducts(input)
	slices.SortStableFunc(ranges, func(a, b Range) int {
		if a.Low < b.Low {
			return -1
		}
		if a.Low > b.Low {
			return 1
		}
		if a.High < b.High {
			return -1
		}
		if a.High > b.High {
			return 1
		}
		return 0
	})
	stage.Println(ranges)
	for {
		oldranges := ranges
		ranges = ReduceRanges(ranges)
		if len(oldranges) == len(ranges) {
			break
		}
	}
	stage.Println(ranges)
	total := 0
	for _, r := range ranges {
		total += r.Len()
	}
	return total, nil
}

func ReduceRanges(ranges []Range) []Range {
	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i].High >= ranges[i+1].Low {
			if ranges[i].High < ranges[i+1].High {
				ranges[i].High = ranges[i+1].High
			}
			ranges[i+1] = Range{}
			i++
			continue
		}
	}
	return slices.DeleteFunc(ranges, func(r Range) bool {
		return r == Range{}
	})
}
