package main

import (
	"testing"

	"github.com/nlm/adventofcode2025/internal/stage"
)

var Stage1TestCases = []stage.TestCase{
	{
		Name:   "example1",
		Result: 21,
	},
	{
		Name:   "input",
		Result: 1660,
	},
}

var Stage2TestCases = []stage.TestCase{
	{
		Name:   "example1",
		Result: 40,
	},
	{
		Name:   "example2",
		Result: 4,
	},
	{
		Name:   "example3",
		Result: 5,
	},
	{
		Name:   "example4",
		Result: 13,
	},
	{
		Name:   "example5",
		Result: 10,
	},
	{
		Name:   "input",
		Result: 305999729392659,
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
