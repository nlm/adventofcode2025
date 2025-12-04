package main

import (
	"testing"

	"github.com/nlm/adventofcode2025/internal/stage"
)

var Stage1TestCases = []stage.TestCase{
	{
		Name:   "example1",
		Result: 1227775554,
	},
	{
		Name:   "example2",
		Result: 1188511885,
	},
	{
		Name:   "input",
		Result: 28846518423,
	},
}

var Stage2TestCases = []stage.TestCase{
	{
		Name:   "example1",
		Result: 4174379265,
	},
	{
		Name:   "example2",
		Result: 1188511885,
	},
	{
		Name:   "input",
		Result: 31578210022,
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
