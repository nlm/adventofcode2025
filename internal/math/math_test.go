package math

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkStdAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int(math.Abs(float64(i)))
		_ = int(math.Abs(float64(-i)))
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Abs(i)
		_ = Abs(-i)
	}
}

func TestMod(t *testing.T) {
	for _, tc := range []struct {
		N   int
		Mod int
		Res int
	}{
		{1, 1, 0},
		{10, 1, 0},
		{1, 2, 1},
		{10, 3, 1},
		{-1, 2, 1},
		{4, -2, 0},
		{-1, -2, 1},
	} {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			assert.Equal(t, tc.Res, Mod(tc.N, tc.Mod))
		})
	}
}
