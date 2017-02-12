#!/bin/bash

set -e

if [[ "$GOPATH" == "" ]]; then
    echo '$GOPATH is empty' 1>&2
    exit 4
fi

LLVM_ORG_DIR="${GOPATH}/src/llvm.org"
LLVM_DIR="${LLVM_ORG_DIR}/llvm"
LLVM_GO_DIR="${LLVM_DIR}/bindings/go"

if [[ -d "$LLVM_DIR" ]]; then
    echo 'LLVM is already installed. Skipped.'
    exit
fi

mkdir -p "$LLVM_ORG_DIR"
cd "$LLVM_ORG_DIR"

git clone --depth 1 -b release_40 --single-branch http://llvm.org/git/llvm.git
cd "$LLVM_GO_DIR"
./build.sh



