#!/usr/bin/env bash

cd project
make docker-build docker-push IMG=gauravagarwalr/k8s-operator-cron:v0.0.1

kind load docker-image gauravagarwalr/k8s-operator-cron:v0.0.1
