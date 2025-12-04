package main

import (
	"testing"

	"github.com/nlm/adventofcode2025/internal/stage"
)

var Stage1TestCases = []stage.TestCase{
	{
		Name:   "example",
		Result: 3,
	},
	{
		Name:   "input",
		Result: 1097,
	},
}

var Stage2TestCases = []stage.TestCase{
	{
		Name:   "example",
		Result: 6,
	},
	{
		Name:   "example2",
		Result: 10,
	},
	{
		Name:   "example3",
		Result: 4,
	},
	{
		Name:   "example4",
		Result: 7,
	},
	{
		Name:   "example5",
		Result: 2,
	},
	{
		Name:   "example6",
		Result: 2,
	},
	{
		Name:   "example7",
		Result: 2,
	},
	{
		Name:   "example8",
		Result: 4,
	},
	{
		Name:   "example9",
		Result: 7,
	},
	{
		Name:   "input",
		Result: 7101,
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
