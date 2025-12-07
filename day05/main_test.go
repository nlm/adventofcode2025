package main

import (
	"testing"

	"github.com/nlm/adventofcode2025/internal/stage"
)

var Stage1TestCases = []stage.TestCase{
	{
		Name:   "example1",
		Result: 3,
	},
	{
		Name:   "input",
		Result: 733,
	},
}

var Stage2TestCases = []stage.TestCase{
	{
		Name:   "example1",
		Result: 14,
	},
	{
		Name:   "input",
		Result: 345821388687084,
	},
}

// Do not edit below

func TestStage1(t *testing.T) {
	stage.Test(t, Stage1, Stage1TestCases)
}

func TestStage2(t *testing.T) {
	stage.Test(t, Stage2, Stage2TestCases)
}

func BenchmarkStage1(b *testing.B) {
	stage.Benchmark(b, Stage1, Stage1TestCases)
}

func BenchmarkStage2(b *testing.B) {
	stage.Benchmark(b, Stage2, Stage2TestCases)
}
