package goop_test

import (
	"github.com/joeycumines/goop/solvers"
	"testing"
)

func TestLPSolve(t *testing.T) {
	t.Run("SimpleMIP", func(t *testing.T) {
		solveSimpleMIPModel(t, solvers.NewLPSolveSolver())
	})

	t.Run("SumRowsCols", func(t *testing.T) {
		solveSumRowsColsModel(t, solvers.NewLPSolveSolver())
	})
}
