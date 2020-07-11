#!/usr/bin/env bash

docker build -t agarwalconsulting/fib-grpc:latest .
docker push agarwalconsulting/fib-grpc:latest
