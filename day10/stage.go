package main

import (
	"io"
	"slices"
	"strings"

	"github.com/nlm/adventofcode2025/internal/combinations"
	"github.com/nlm/adventofcode2025/internal/iterators"
	"github.com/nlm/adventofcode2025/internal/matrix"
	"github.com/nlm/adventofcode2025/internal/solver"
	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

type Machine struct {
	Lights   []bool
	Wirings  [][]int
	Joltages []int
}

func (m *Machine) PushButton(lights []bool, button int) {
	for _, v := range m.Wirings[button] {
		lights[v] = !lights[v]
	}
}

func (m *Machine) PushButtonJoltages(joltages []int, button int) {
	for _, v := range m.Wirings[button] {
		joltages[v]++
	}
}

func (m Machine) LightsMatch(lights []bool) bool {
	for i := range m.Lights {
		if m.Lights[i] != lights[i] {
			return false
		}
	}
	return true
}

func (m Machine) JoltagesMatch(joltages []int) bool {
	for i := range m.Joltages {
		if m.Joltages[i] != joltages[i] {
			return false
		}
	}
	return true
}

func (m Machine) OverJoltages(joltages []int) bool {
	for i := range m.Joltages {
		if m.Joltages[i] < joltages[i] {
			return true
		}
	}
	return false
}

func (m Machine) NewLights() []bool {
	return make([]bool, len(m.Lights))
}

func (m Machine) NewJoltages() []int {
	return make([]int, len(m.Joltages))
}

func ParseMachine(line string) *Machine {
	m := Machine{}
	parts := strings.Split(line, " ")
	// Lights
	lights := strings.Trim(parts[0], "[]")
	for _, c := range lights {
		m.Lights = append(m.Lights, c == '#')
	}
	// Wiring
	for _, part := range parts[1 : len(parts)-1] {
		part = strings.Trim(part, "()")
		wiring := make([]int, 0)
		for _, s := range strings.Split(part, ",") {
			wiring = append(wiring, utils.MustAtoi(s))
		}
		m.Wirings = append(m.Wirings, wiring)
	}
	// Joltage
	joltages := strings.Trim(parts[len(parts)-1], "{}")
	for _, s := range strings.Split(joltages, ",") {
		m.Joltages = append(m.Joltages, utils.MustAtoi(s))
	}
	return &m
}

func HandleMachineLights(m *Machine, maxK int) int {
	stage.Println("machine:", *m)
	idxs := slices.Collect(iterators.Range(len(m.Wirings)))
	stage.Println("presses:", idxs)
	stage.Println("lights:", m.Lights)
	for k := 1; k < maxK+1; k++ {
		stage.Println("\nk:", k)
		for presses := range combinations.CartesianProduct(idxs, k) {
			lights := m.NewLights()
			stage.Println("presses:", presses)
			for _, button := range presses {
				m.PushButton(lights, button)
				stage.Println(lights)
				if m.LightsMatch(lights) {
					stage.Println("EUREKA", presses)
					return k
				}
			}
		}
	}
	return -1
}

// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
//
// a [0, 0, 0, 1] + b [0, 1, 0, 1] + c [0, 0, 1, 0] + d [0, 0, 1, 1] + e [1, 0, 0, 1] + f [1, 1, 0, 0] = [3, 5, 4, 7]
// e + f = 3
// b + f = 5
// c + d = 4
// a + b + d + e = 7
func HandleMachineJoltages2(m *Machine) int {
	// prepare equation matrix
	// set variables indices
	mx := matrix.New[int](len(m.Wirings)+1, len(m.Joltages))
	for i, w := range m.Wirings {
		for _, idx := range w {
			mx.SetAt(i, idx, 1)
		}
	}
	// set results
	for i, j := range m.Joltages {
		mx.SetAt(mx.Size.X-1, i, j)
	}

	stage.Println(matrix.IMatrix(mx))
	res := utils.Must(solver.Solve(mx))
	stage.Println("res:", res)
	total := 0
	for _, n := range res {
		total += n
	}
	return total
}

func HandleMachineJoltages(m *Machine, maxK int) int {
	stage.Println("machine:", *m)
	idxs := slices.Collect(iterators.Range(len(m.Wirings)))
	stage.Println("presses:", idxs)
	stage.Println("joltages:", m.Joltages)
	iter := 0
	for k := 1; k < maxK+1; k++ {
		// fmt.Println("\nk:", k)
		for presses := range combinations.CartesianProduct(idxs, k) {
			joltages := m.NewJoltages()
			// stage.Println("presses:", presses)
			for _, button := range presses {
				iter++
				m.PushButtonJoltages(joltages, button)
				// stage.Println(joltages)
				if m.OverJoltages(joltages) {
					break
				}
				if m.JoltagesMatch(joltages) {
					// stage.Println("EUREKA", presses)
					return k
				}
			}
		}
		// fmt.Println(iter)
	}
	return -1
}

func ParseMachines(input io.Reader) []*Machine {
	machines := make([]*Machine, 0)
	for line := range iterators.MustLines(input) {
		stage.Println("parse:", line)
		machine := ParseMachine(line)
		machines = append(machines, machine)
	}
	return machines
}

func Stage1(input io.Reader) (any, error) {
	const maxK = 10
	res := 0
	machines := ParseMachines(input)
	for _, machine := range machines {
		np := HandleMachineLights(machine, maxK)
		stage.Println(">> FOUND", np, "for", *machine)
		res += np
	}
	return res, nil
}

func Stage2(input io.Reader) (any, error) {
	res := 0
	machines := ParseMachines(input)
	for _, machine := range machines {
		np := HandleMachineJoltages2(machine)
		stage.Println(">> FOUND", np, "for", *machine)
		res += np
	}
	return res, nil
}
