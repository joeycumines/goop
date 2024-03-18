package solvers

// #cgo CXXFLAGS: --std=c++11 -I/Library/gurobi903/mac64/include -I/Users/joeyc/dev/goop/include
// #cgo CXXFLAGS: -I/opt/homebrew/opt/lp_solve/include
// #cgo LDFLAGS: -L/Library/gurobi903/mac64/lib -lgurobi_c++ -lgurobi90 
// #cgo LDFLAGS: -L/opt/homebrew/opt/lp_solve/lib -llpsolve55
 import "C"
