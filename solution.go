package goop

import (
	"github.com/joeycumines/goop/internal/solvers"
)

const (
	tinyNum float64 = 0.01
)

// Solution stores the solution of an optimization problem and associated
// metadata.
type Solution struct {
	// variable values
	vals []float64

	// The objective for the solution
	Objective float64

	// Whether or not the solution is within the optimality threshold
	Optimal bool

	// The optimality gap returned from the solver. For many solvers, this is
	// the gap between the best possible solution with integer relaxation and
	// the best integer solution found so far.
	Gap float64
}

// newSolution makes a memory-safe copy of the solution data. Note that the
// input solution MUST be freed by the caller, using the appropriate
// (swig-generated) delete function.
func newSolution(mipSol solvers.MIPSolution) *Solution {
	vec := mipSol.GetValues()
	vals := make([]float64, int(vec.Size()))
	for i := range vals {
		vals[i] = vec.Get(i)
	}
	return &Solution{
		vals:      vals,
		Objective: mipSol.GetObj(),
		Optimal:   mipSol.GetOptimal(),
		Gap:       mipSol.GetGap(),
	}
}

// Value returns the value assigned to the variable in the solution
func (s *Solution) Value(v *Var) float64 {
	return s.vals[int(v.ID())]
}

// IsOne returns true if the value assigned to the variable is an integer,
// and assigned to one. This is a convenience method which should not be
// super trusted...
func (s *Solution) IsOne(v *Var) bool {
	return (v.Type() == Integer || v.Type() == Binary) && s.Value(v) > tinyNum
}
