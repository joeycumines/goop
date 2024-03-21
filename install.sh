#!/usr/bin/env bash

set -xe

if [[ $(uname) == 'Linux' ]]; then
    sudo apt install -y libpcre3 libpcre3-dev autotools-dev byacc flex cmake build-essential autoconf lp-solve liblpsolve55-dev swig
elif [[ $(uname) == 'Darwin' ]]; then
    brew install pcre autoconf
    brew install swig
    brew install lp_solve
    mkdir -p include
    for f in /opt/homebrew/opt/lp_solve/include/*; do ln -snf "$f" include/"$(basename "$f")"; done
    mkdir -p lib
    for f in /opt/homebrew/opt/lp_solve/lib/*; do ln -snf "$f" lib/"$(basename "$f")"; done
fi
