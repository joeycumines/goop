#!/usr/bin/env bash

set -xe

if [[ $(uname) == 'Linux' ]]; then
    sudo apt install -y libpcre3 libpcre3-dev autotools-dev byacc flex cmake build-essential autoconf lp-solve liblpsolve55-dev swig
elif [[ $(uname) == 'Darwin' ]]; then
    brew install pcre autoconf
    brew install swig
    brew install lp_solve
fi
