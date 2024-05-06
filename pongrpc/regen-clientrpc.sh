#!/bin/sh

BINDIR=$(mktemp -d)

build_protoc_gen_go() {
    mkdir -p $BINDIR
    export GOBIN=$BINDIR
}

generate() {
    protoc --go_out=. --go-grpc_out=. pong.proto
    protoc -I=. ./pong.proto --js_out=import_style=commonjs,binary:../ui-client/src/pongrpc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../ui-client/src/pongrpc
}

# Build the bins from the main module, so that clientrpc doesn't need to
# require all client and tool dependencies.
(cd .. && build_protoc_gen_go)
GENPATH="$BINDIR:$PATH"
PATH=$GENPATH generate
