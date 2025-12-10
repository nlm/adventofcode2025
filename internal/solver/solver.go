package solver

import (
	"errors"
	"fmt"

	"github.com/aclements/go-z3/z3"
	"github.com/nlm/adventofcode2025/internal/matrix"
)

var (
	ErrNoSolution = errors.New("solver: no solution found")
)

func z3Int(ctx *z3.Context, i int) z3.Int {
	return ctx.FromInt(int64(i), ctx.IntSort()).(z3.Int)
}

// Solves a linear equation represented by a matrix
//
// example:
// a1x + b1y + c1z = d1
// a2x + b2y + c2z = d2
//
// matrix representation:
// [a1, b1, c1, d1]
// [a2, b2, c2, d2]
//
// it will return a slice of int:
// [a, b, c]
//
// example code:
//
//	func main() {
//		m := utils.Must(matrix.NewFromSeq(7, 4, slices.Values([]int{
//			0, 0, 0, 0, 1, 1, 3,
//			0, 1, 0, 0, 0, 1, 5,
//			0, 0, 1, 1, 0, 0, 4,
//			1, 1, 0, 1, 1, 0, 7,
//		})))
//		fmt.Println(matrix.IMatrix(m))
//		fmt.Println(utils.Must(solver.Solve(m)))
//	}
func Solve(m *matrix.Matrix[int]) ([]int, error) {
	ctx := z3.NewContext(nil)
	slv := z3.NewSolver(ctx)
	zero := z3Int(ctx, 0)

	// prepare variables, last col is result
	vars := make([]z3.Int, m.Size.X-1)
	for i := range m.Size.X - 1 {
		v := ctx.IntConst(fmt.Sprintf("x%d", i))
		vars[i] = v
		slv.Assert(v.GE(zero))
	}

	// each equation
	for y := range m.Size.Y {
		// caculate components
		items := make([]z3.Int, 0)
		for x := range m.Size.X - 1 {
			items = append(items, vars[x].Mul(z3Int(ctx, m.At(x, y))))
		}
		// add constraint to solver
		slv.Assert(zero.Add(items...).Eq(z3Int(ctx, m.At(m.Size.X-1, y))))
	}

	sat, err := slv.Check()
	if err != nil {
		return nil, fmt.Errorf("solver error: %w", err)
	}

	if !sat {
		return nil, ErrNoSolution
	}

	model := slv.Model()

	// FIXME: Custom Code to find lowest sum
	sum := zero.Add(vars...)
	best := model.Eval(sum, true)
	for {
		slv.Push()
		slv.Assert(sum.LT(best.(z3.Int)))
		sat, err := slv.Check()
		if err != nil {
			return nil, fmt.Errorf("solver error: %w", err)
		}
		if !sat {
			slv.Pop()
			break
		}
		model = slv.Model()
		best = model.Eval(sum, true)
	}
	// End

	res := make([]int, 0, m.Size.X-1)
	for i := range vars {
		v, _, _ := model.Eval(vars[i], true).(z3.Int).AsInt64()
		res = append(res, int(v))
	}

	return res, nil
}
