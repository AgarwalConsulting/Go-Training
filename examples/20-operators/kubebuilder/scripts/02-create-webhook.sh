#!/usr/bin/env bash

cd project

kubebuilder create webhook --group batch --version v1 --kind CronJob --defaulting --programmatic-validation
