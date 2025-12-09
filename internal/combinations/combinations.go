package combinations

import (
	"iter"
	"slices"
)

// cartesianProductIndexes generates all the possible k-tuples
// from a set of size n. items are present more than once.
//
// Internally it's using an array of indexes to iterate over every possible
// combination of positions in the elts list. It's using basic math to pass to
// the next possiblity by adding 1 to the lowest weight index and propagating
// the carry to a higher index, until all solutions have been interated over.
func cartesianProductIndexes(n, k int) iter.Seq[[]int] {
	if n < 1 {
		panic("n must be >1")
	}
	if k < 0 {
		panic("k must be >=0")
	}
	return func(yield func([]int) bool) {
		indexes := make([]int, k)
		for {
			// return current combination
			if !yield(slices.Clone(indexes)) {
				return
			}
			// calculate indexes for next combination
			carry := 0
			for i := k - 1; i >= 0; i-- {
				indexes[i]++
				if indexes[i] >= n {
					carry++
					indexes[i] = 0
				} else {
					break
				}
			}
			// return if every bit triggered carry
			if carry == k {
				return
			}
		}
	}
}

// CartesianProduct iterates over every combination of n items from the elts list.
// It will allocate "1 + len(elts) to the power of n" slices.
func CartesianProduct[T any](elts []T, k int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(elts) == 0 || k <= 0 || k > len(elts) {
			return
		}
		for indexes := range cartesianProductIndexes(len(elts), k) {
			data := make([]T, k)
			for i := range k {
				data[i] = elts[indexes[i]]
			}
			if !yield(data) {
				return
			}
		}
	}
}

// combinationsIndexes generates all the possible unique k-sets
// from a set of size n. items are present only once.
// This version is not optimized, as it generates all the
// possibilities and filters out bad generations.
func combinationsIndexes(n, k int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
	outer:
		for indexes := range cartesianProductIndexes(n, k) {
			maxVal := indexes[0]
			// optimization to cut the tail
			if maxVal == n-1 {
				return
			}
			for i := 1; i < len(indexes); i++ {
				if indexes[i] <= maxVal {
					continue outer
				}
			}
			if !yield(indexes) {
				return
			}
		}
	}
}

// Combinations generates all the possible unique k-sets
// of items from elts. items are present only once.
func Combinations[T any](elts []T, k int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(elts) == 0 || k <= 0 || k > len(elts) {
			// panic instead ?
			return
		}
		for indexes := range combinationsIndexes(len(elts), k) {
			data := make([]T, k)
			for i := range k {
				data[i] = elts[indexes[i]]
			}
			if !yield(data) {
				return
			}
		}
	}
}
