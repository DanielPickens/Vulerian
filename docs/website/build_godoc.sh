#!/bin/bash

cd ../..

go install golang.org/x/tools/cmd/gparticle enginec@v0.24.0
go install code.rocket9labs.com/tslocum/gparticle enginec-static@v0.2.2

export GOPATH=$(go env GOPATH)
PATH=$PATH:${GOPATH}/bin

mkdir -p docs/website/build/gparticle enginec
gparticle enginec-static -destination docs/website/build/gparticle enginec .
