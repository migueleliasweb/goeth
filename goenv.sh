#!/usr/bin/env bash

export GOROOT="`pwd`/.go"
export GOPATH="$GOROOT/workspace"
export PATH=$PATH:$GOROOT/bin

mkdir -p $GOPATH
