#!/usr/bin/env bash

# https://tyk.io/docs/tyk-oss/ce-docker/

set -e

docker network create tyk

docker run -itd --rm --name tyk-redis --network tyk -p 127.0.0.1:6379:6379 redis:4.0-alpine

# wget https://raw.githubusercontent.com/TykTechnologies/tyk-gateway-docker/master/tyk.standalone.conf

sleep 10

docker run \
  --rm \
  --name tyk_gateway \
  --network tyk \
  -p 8080:8080 \
  -v $(pwd)/tyk.standalone.conf:/opt/tyk-gateway/tyk.conf \
  -v $(pwd)/apps:/opt/tyk-gateway/apps \
  docker.tyk.io/tyk-gateway/tyk-gateway:latest
