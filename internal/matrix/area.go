package matrix

import "iter"

type Coorder[T comparable] interface {
	InCoord(c Coord) bool
	AtCoord(c Coord) T
	Coords() iter.Seq[Coord]
}

var _ Coorder[struct{}] = (*Area[struct{}])(nil)

type Area[T comparable] struct {
	m      *Matrix[T]
	origin Coord
	size   Vec
	max    Coord
}

func NewArea[T comparable](m *Matrix[T], origin Coord, size Vec) *Area[T] {
	if origin.X+m.Size.X > size.X || origin.Y+m.Size.Y > size.Y {
		panic("Area out of bounds")
	}
	return &Area[T]{
		m:      m,
		origin: origin,
		size:   size,
		max:    origin.Add(size).Add(Vec{-1, -1}),
	}
}

func (a Area[T]) Size() Vec {
	return a.size
}

func (a *Area[T]) InCoord(c Coord) bool {
	return c.X >= a.origin.X && c.X <= a.max.X && c.Y >= a.origin.Y && c.Y <= a.max.Y
}

func (a *Area[T]) AtCoord(c Coord) T {
	return a.m.AtCoord(c.Add(Vec(a.origin)))
}

// Coords will return an iterator over all the coordinates
// that exist in the matrix.
func (a *Area[T]) Coords() iter.Seq[Coord] {
	return func(yield func(Coord) bool) {
		for y := range a.size.Y {
			for x := range a.size.X {
				if !yield(Coord{X: x, Y: y}) {
					return
				}
			}
		}
	}
}
