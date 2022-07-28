#!/usr/bin/env bash

cd project
make docker-build docker-push IMG=gauravagarwalr/k8s-operator-cron:v0.0.1
