package math

import (
	"golang.org/x/exp/constraints"
)

// Number is a type constraint matching Integers and Floats.
type Number interface {
	constraints.Integer | constraints.Float
}

// Abs returns the absolute value of n.
func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// Mod returns the positive modulus of 'n' by 'mod'.
func Mod[T constraints.Integer](n T, mod T) T {
	n %= mod
	switch {
	case n >= 0:
		return n
	case mod < 0:
		return n - mod
	default:
		return n + mod
	}
}
