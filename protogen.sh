#!/bin/bash

HERE=$(dirname $0)
PROTO_DIR=${HERE}/proto
GO_OUT_DIR=${HERE}/_proto

mkdir -p ${GO_OUT_DIR}

if [[ "$GOBIN" == "" ]]; then
  if [[ "$GOPATH" == "" ]]; then
    echo "Required env var GOPATH is not set; aborting with error; see the following documentation which can be invoked via the 'go help gopath' command."
    go help gopath
    exit -1
  fi

  echo "Optional env var GOBIN is not set; using default derived from GOPATH as: \"$GOPATH/bin\""
  export GOBIN="$GOPATH/bin"
fi

echo "Compiling protobuf definitions"
protoc \
  --plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
  -I ${PROTO_DIR} \
  --go_out=plugins=grpc:${GO_OUT_DIR} \
  ${PROTO_DIR}/osechet/test/test_types.proto
