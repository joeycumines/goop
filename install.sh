#!/usr/bin/env bash

set -e

WORK=$(pwd)

if [[ $(uname) == 'Linux' ]]; then
    sudo apt install -y libpcre3 libpcre3-dev autotools-dev byacc \
        flex cmake build-essential autoconf
    scripts/install_swig.sh
    scripts/install_lpsolve.sh
    mkdir -p .third_party
elif [[ $(uname) == 'Darwin' ]]; then
    brew install pcre autoconf
    brew install swig 
    brew install lp_solve 

    go run scripts/make_lib.go --go-fname internal/solvers/lib.go --pkg solvers
fi
