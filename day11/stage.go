package main

import (
	"fmt"
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

type Memo map[string]int

func (m Memo) FindPath(from, to string, devices map[string]Device) int {
	total := 0
	for _, out := range devices[from].Outputs {
		switch out {
		case to:
			total++
		default:
			v, ok := m[out]
			if !ok {
				v = m.FindPath(out, to, devices)
				m[out] = v
			}
			total += v
		}
	}
	return total
}

func Stage1(input io.Reader) (any, error) {
	devices := make(map[string]Device, 0)
	for line := range iterators.MustLines(input) {
		device := ParseDevice(line)
		devices[device.Name] = device
	}
	m := make(Memo, 0)
	res := m.FindPath("you", "out", devices)
	return res, nil
}

func Stage2(input io.Reader) (any, error) {
	devices := make(map[string]Device, 0)
	for line := range iterators.MustLines(input) {
		device := ParseDevice(line)
		devices[device.Name] = device
	}

	m := make(Memo, 0)
	clear(m)
	fftDac := m.FindPath("fft", "dac", devices)
	if fftDac != 0 {
		clear(m)
		svrFft := m.FindPath("svr", "fft", devices)
		clear(m)
		dacOut := m.FindPath("dac", "out", devices)
		return svrFft * fftDac * dacOut, nil
	}

	clear(m)
	dacFft := m.FindPath("dac", "fft", devices)
	if dacFft != 0 {
		clear(m)
		svrDac := m.FindPath("svr", "dac", devices)
		clear(m)
		fftOut := m.FindPath("fft", "out", devices)
		return svrDac * dacFft * fftOut, nil
	}

	return 0, fmt.Errorf("no solution found")
}
