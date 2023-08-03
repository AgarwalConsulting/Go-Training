#!/usr/bin/env bash

go test -v -bench=BenchmarkConcurrentFrequencyWO --benchmem ./... -cpuprofile=cpuwo.prof -memprofile=memwo.prof -trace=tracewo.out

go test -v -bench=BenchmarkConcurrentFrequencyWC --benchmem ./... -cpuprofile=cpuwc.prof -memprofile=memwc.prof -trace=tracewc.out
