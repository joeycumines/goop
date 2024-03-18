package goop_test

import (
	"testing"
)

func TestLPSolve(t *testing.T) {
	t.Run("SimpleMIP", func(t *testing.T) {
		solveSimpleMIPModel(t)
	})

	t.Run("SumRowsCols", func(t *testing.T) {
		solveSumRowsColsModel(t)
	})
}
