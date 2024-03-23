package solvers

// #cgo CXXFLAGS: --std=c++11 -I/usr/local/include/lpsolve -I/opt/homebrew/opt/lp_solve/include -I/usr/include/lpsolve
// #cgo LDFLAGS: -llpsolve55 -L/usr/local/lib -L/opt/homebrew/opt/lp_solve/lib -L/usr/lib -L/usr/lib64
import "C"

// TODO: Add support for replacing the paths above and, instead using a single path each, specified by env var, behind a build tag.
