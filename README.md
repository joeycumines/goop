# Goop [![Go Report Card](https://goreportcard.com/badge/github.com/joeycumines/goop)](https://goreportcard.com/report/github.com/joeycumines/goop) [![Build Status](https://travis-ci.org/mit-drl/goop.svg?branch=master)](https://travis-ci.org/mit-drl/goop) [![Go Doc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=round-square)](https://godoc.org/github.com/joeycumines/goop) [![Maintainability](https://api.codeclimate.com/v1/badges/7bb0cc28fd6d18d2de44/maintainability)](https://codeclimate.com/github/mit-drl/goop/maintainability) [![codecov](https://codecov.io/gh/mit-drl/goop/branch/master/graph/badge.svg)](https://codecov.io/gh/mit-drl/goop)

General Linear Optimization in Go. `goop` provides general interface for solving
mixed integer linear optimization problems using a variety of back-end solvers
including LPSolve and Gurobi.

# Quickstart

We are going to start with a simple example showing how `goop` can be used to
solve integer linear programs. The example below seeks to maximize the following
MIP:

```
maximize    x +   y + 2 z
subject to  x + 2 y + 3 z <= 4
            x +   y       >= 1
x, y, z binary
```

This is is the same example implemented [here](http://www.gurobi.com/documentation/7.5/examples/mip1_py.html). Below
we have implemented the model using `goop` and have optimized the model using
the supported Gurobi solver.

```go
package main

import (
	"fmt"
	"github.com/joeycumines/goop"
	"github.com/joeycumines/goop/internal/solvers"
)

func main() {
	// Instantiate a new model
	m := goop.NewModel()

	// Add your variables to the model
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()
	z := m.AddBinaryVar()

	// Add your constraints
	m.AddConstr(goop.Sum(x, y.Mult(2), z.Mult(3)).LessEq(goop.K(4)))
	m.AddConstr(goop.Sum(x, y).GreaterEq(goop.One))

	// Set a linear objective using your variables
	obj := goop.Sum(x, y, z.Mult(2))
	m.SetObjective(obj, goop.SenseMaximize)

	// Optimize the variables according to the model
	sol, err := m.Optimize(solvers.NewGurobiSolver())

	// Check if there is an error from the solver. No error should be returned
	// for this model
	if err != nil {
		panic("Should not have an error")
	}

	// Print out the solution
	fmt.Println("x =", sol.Value(x))
	fmt.Println("y =", sol.Value(y))
	fmt.Println("z =", sol.Value(z))

	// Output:
	// x = 1
	// y = 0
	// z = 1
}
```

# Installation

1. First get the code
```
mkdir -p $GOPATH/github.com/mit-drl && cd $GOPATH/github.com/mit-drl
git clone https://github.com/joeycumines/goop && cd goop
```

2. Next build install the dependencies
```
./install.sh
```

3. Follow the [instructions](#Solver Notes) for your solver of choice. Note,
currently only Gurobi is supported

4. Finally build the library
```
go build
```
Note that due to a quirk with Gurobi, if you are using Ubuntu < 16.04, you must
build with
```
go build -tags pre_xenial
```

5. (Optional) Test our installation
```
govendor test -v +local
```

# Solver Notes

We currently have bindings for Gurobi and LPSolve. Please follow the
instructions below for using these specific solvers.

## Gurobi
- You must have [Gurobi](http://www.gurobi.com/downloads/download-center)
installed and have a valid license.
- The `GUROBI_HOME` environment variable must be set to the home directory
of your Gurobi installation

## LPSolve
LPSolve is installed using the normal install procedure and should work out of
the box.
