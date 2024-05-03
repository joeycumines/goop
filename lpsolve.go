package goop

const (
	// Scaling options

	ScaleNone       ScaleMode = 0 // No scaling is applied.
	ScaleExtreme    ScaleMode = 1 // Scale to convergence using largest absolute value.
	ScaleRange      ScaleMode = 2 // Scale based on the simple numerical range.
	ScaleMean       ScaleMode = 3 // Numerical range-based scaling.
	ScaleGeometric  ScaleMode = 4 // Geometric scaling.
	ScaleFuture1    ScaleMode = 5 // Reserved for future use.
	ScaleFuture2    ScaleMode = 6 // Reserved for future use.
	ScaleCurtisReid ScaleMode = 7 // Curtis-Reid "optimal" scaling.

	// Alternative scaling weights

	ScaleLinear      ScaleMode = 0  // Linear scaling.
	ScaleQuadratic   ScaleMode = 8  // Quadratic scaling.
	ScaleLogarithmic ScaleMode = 16 // Scale using logarithmic mean of all values.
	ScaleUserWeight  ScaleMode = 31 // User-specified weights.

	// Scaling modes

	ScalePower2      ScaleMode = 32   // Scale rounded to power of 2.
	ScaleEquilibrate ScaleMode = 64   // Ensure no scaled number is above 1.
	ScaleIntegers    ScaleMode = 128  // Apply scaling to integer columns/variables.
	ScaleDynUpdate   ScaleMode = 256  // Incrementally apply scaling for every solve.
	ScaleRowsOnly    ScaleMode = 512  // Scale only the rows.
	ScaleColsOnly    ScaleMode = 1024 // Scale only the columns.

	// Predefined scaling models

	ScaleModelEquilibrated ScaleMode = ScaleLinear + ScaleExtreme + ScaleIntegers
	ScaleModelGeometric    ScaleMode = ScaleLinear + ScaleGeometric + ScaleIntegers
	ScaleModelArithmetic   ScaleMode = ScaleLinear + ScaleMean + ScaleIntegers
	ScaleModelDynamic      ScaleMode = ScaleModelGeometric + ScaleEquilibrate
	ScaleModelCurtisReid   ScaleMode = ScaleCurtisReid + ScaleIntegers + ScalePower2
)

type (
	// ScaleMode models lp_solve scaling options.
	// See also [Model.SetScaling].
	ScaleMode int
)

// SetScaling sets the scaling mode for a lp_solve model.
// This function configures the scaling algorithm to be used in the solver to
// enhance numerical stability and performance. The specified scaling mode can
// significantly influence the solution process by improving accuracy and
// reducing computation time.
//
// Scaling should be applied before the solve function and can include a
// combination of basic scaling types and additional modes. It's advisable to
// enable scaling to prevent numerical issues, especially when dealing with
// data of varying magnitudes.
//
// Parameters:
//
//   - scaleMode: Specifies the combination of scaling options to use. Can
//     include basic scaling types such as ScaleExtreme or ScaleGeometric
//     combined using bitwise OR with modes like ScaleEquilibrate or
//     ScaleIntegers to tailor the scaling process to specific needs.
//
// Example usage:
//
//	SetScaling(lp, ScaleModelGeometric | ScaleDynUpdate)
//
// See also https://lpsolve.sourceforge.net/5.5/set_scaling.htm
func (m *Model) SetScaling(scaleMode ScaleMode) {
	m.scaleMode = &scaleMode
}
