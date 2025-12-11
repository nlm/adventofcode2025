package main

import (
	"io"
	"strings"

	"github.com/nlm/adventofcode2025/internal/iterators"
)

type Device struct {
	Name    string
	Outputs []string
}

func ParseDevice(line string) Device {
	line = strings.ReplaceAll(line, "  ", " ")
	line = strings.Replace(line, ":", "", 1)
	items := strings.Split(line, " ")
	return Device{
		Name:    items[0],
		Outputs: items[1:],
	}
}

var Memo = make(map[string]int)

func Stage1(input io.Reader) (any, error) {
	devices := make(map[string]Device, 0)
	for line := range iterators.MustLines(input) {
		device := ParseDevice(line)
		devices[device.Name] = device
	}
	clear(Memo)
	res := FindPath("you", "out", "", devices)
	return res, nil
}

func FindPath(from, to, avoid string, devices map[string]Device) int {
	total := 0
	for _, out := range devices[from].Outputs {
		switch out {
		case to:
			total++
		case avoid:
			return 0
		default:
			v, ok := Memo[out]
			if !ok {
				v = FindPath(out, to, avoid, devices)
				Memo[out] = v
			}
			total += v
		}
	}
	return total
}

func Stage2(input io.Reader) (any, error) {
	devices := make(map[string]Device, 0)
	for line := range iterators.MustLines(input) {
		device := ParseDevice(line)
		devices[device.Name] = device
	}
	clear(Memo)
	dacOut := FindPath("dac", "out", "fft", devices)
	clear(Memo)
	fftDac := FindPath("fft", "dac", "", devices)
	clear(Memo)
	fftOut := FindPath("fft", "out", "dac", devices)
	clear(Memo)
	dacFft := FindPath("dac", "fft", "", devices)
	clear(Memo)
	svrFft := FindPath("svr", "fft", "dac", devices)
	clear(Memo)
	svrDac := FindPath("svr", "dac", "fft", devices)

	return svrDac*dacFft*fftOut + svrFft*fftDac*dacOut, nil
}
