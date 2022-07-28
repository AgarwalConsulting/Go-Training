#!/usr/bin/env bash

cd project

kubebuilder create api --group batch --version v1 --kind CronJob
