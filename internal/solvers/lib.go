package solvers

// #cgo CXXFLAGS: --std=c++11 -I../../include -I/usr/include/lpsolve
// #cgo LDFLAGS: -L../../lib -L/usr/lib/lp_solve -L/usr/lib64 -llpsolve55
import "C"
