package main

import (
	"io"
	"strings"

	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/matrix"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

func Stage1(input io.Reader) (any, error) {
	intItems := make([][]int, 0)
	opItems := make([]string, 0)
	for line := range iterators.MustLines(input) {
		// intItems = append(intItems, []int{})
		stage.Println("newline")
		idx := 0
		for _, item := range strings.Split(line, " ") {
			item := strings.TrimSpace(item)
			switch item {
			case "":
				continue
			case "+", "*":
				stage.Printf("operator: '%s'\n", item)
				opItems = append(opItems, item)
			default:
				val := utils.MustAtoi(item)
				stage.Printf("number: '%d'\n", val)
				if len(intItems) < idx+1 {
					intItems = append(intItems, []int{})
				}
				intItems[idx] = append(intItems[idx], val)
			}
			idx++
		}
		stage.Println(intItems)
	}
	stage.Println(opItems)
	total := 0
	for i := range opItems {
		switch opItems[i] {
		case "+":
			total += iterators.ReduceSlice(0, intItems[i], func(a, b int) int { return a + b })
		case "*":
			total += iterators.ReduceSlice(1, intItems[i], func(a, b int) int { return a * b })
		}
	}
	return total, nil
}

func Stage2(input io.Reader) (any, error) {
	m := utils.Must(matrix.NewFromReader(input))
	stage.Print(matrix.SMatrix(m))
	total := 0
	stack := make([]int, 0)
	for x := m.Size.X - 1; x >= 0; x-- {
		v := 0
		defined := false
		for y := 0; y < m.Size.Y; y++ {
			char := m.At(x, y)
			stage.Println(string(char))
			switch char {
			case ' ':
				if defined {
					stage.Println("stack", v)
					stack = append(stack, v)
					v = 0
					defined = false
				}
			case '+':
				if defined {
					stage.Println("stack", v)
					stack = append(stack, v)
					v = 0
					defined = false
				}
				stage.Println("add", stack)
				total += iterators.ReduceSlice(0, stack, func(a, b int) int { return a + b })
				stack = stack[:0]
			case '*':
				if defined {
					stage.Println("stack", v)
					stack = append(stack, v)
					v = 0
					defined = false
				}
				stage.Println("mul", stack)
				total += iterators.ReduceSlice(1, stack, func(a, b int) int { return a * b })
				stack = stack[:0]
			default:
				stage.Println("read", utils.MustAtoi(string(char)))
				v = v*10 + utils.MustAtoi(string(char))
				defined = true
			}
		}
	}
	return total, nil
}
