package iterators

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	for _, tc := range []struct {
		Slice  []int
		Result int
	}{
		{[]int{1, 2, 3, 4}, 10},
	} {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			assert.Equal(t, tc.Result, Reduce(0, slices.Values(tc.Slice), func(a, b int) int { return a + b }))
		})
	}
}
