#!/usr/bin/env bash

docker run -it --rm \
  -e GRPC_PORT=$GRPC_PORT \
  -e DIAGNOSTICS_PORT=$DIAGNOSTICS_PORT \
  -p $GRPC_PORT:$GRPC_PORT \
  -p $DIAGNOSTICS_PORT:$DIAGNOSTICS_PORT \
  agarwalconsulting/fib-grpc:latest
