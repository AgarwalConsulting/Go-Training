#!/usr/bin/env bash

# See: https://book.kubebuilder.io/multiversion-tutorial/conversion-concepts.html

cd project
kubebuilder create api --group batch --version v2 --kind CronJob
