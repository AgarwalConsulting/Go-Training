#!/usr/bin/env bash

set -e

mkdir tyk-plugin
cd tyk-plugin
go mod init tyk-plugin
go get github.com/TykTechnologies/tyk@ffa83a27d3bf793aa27e5f6e4c7106106286699d
go mod tidy


# Building the plugin
# docker pull tykio/tyk-plugin-compiler:v5.2.1
# docker run --rm -v `pwd`:/plugin-source \
#            --platform=linux/amd64 \
#            tykio/tyk-plugin-compiler:v5.2.1 plugin.so

# Loading the plugin: https://tyk.io/docs/plugins/supported-languages/golang/#loading-the-plugin
# Updating the plugin: https://tyk.io/docs/plugins/supported-languages/golang/#updating-the-plugin

# Plugin Types: https://tyk.io/docs/plugins/supported-languages/golang/#plugin-types
