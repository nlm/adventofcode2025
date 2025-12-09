package combinations

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartesianProduct(t *testing.T) {
	for _, tc := range []struct {
		Elts []string
		K    int
		Res  [][]string
	}{
		{[]string{}, 0, nil},
		{[]string{}, 2, nil},
		{[]string{"A", "B", "C"}, 4, nil},
		{[]string{"A", "B"}, 2, [][]string{{"A", "A"}, {"A", "B"}, {"B", "A"}, {"B", "B"}}},
		{[]string{"A", "B", "C"}, 2, [][]string{
			{"A", "A"}, {"A", "B"}, {"A", "C"},
			{"B", "A"}, {"B", "B"}, {"B", "C"},
			{"C", "A"}, {"C", "B"}, {"C", "C"},
		}},
	} {
		t.Run(fmt.Sprint(tc.Elts, " ", tc.K), func(t *testing.T) {
			res := slices.Collect(CartesianProduct(tc.Elts, tc.K))
			assert.Equal(t, tc.Res, res)
		})
	}

}

func TestCombinations(t *testing.T) {
	for _, tc := range []struct {
		Elts []string
		K    int
		Res  [][]string
	}{
		{[]string{}, 0, nil},
		{[]string{}, 2, nil},
		{[]string{"A", "B", "C"}, 4, nil},
		{[]string{"A", "B"}, 2, [][]string{{"A", "B"}}},
		{[]string{"A", "B", "C"}, 2, [][]string{{"A", "B"}, {"A", "C"}, {"B", "C"}}},
		{[]string{"A", "B", "C", "D"}, 2, [][]string{{"A", "B"}, {"A", "C"}, {"A", "D"}, {"B", "C"}, {"B", "D"}, {"C", "D"}}},
	} {
		t.Run(fmt.Sprint(tc.Elts, " ", tc.K), func(t *testing.T) {
			res := slices.Collect(Combinations(tc.Elts, tc.K))
			assert.Equal(t, tc.Res, res)
		})
	}

}
