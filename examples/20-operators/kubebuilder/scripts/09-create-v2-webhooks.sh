#!/usr/bin/env bash

kubebuilder create webhook --group batch --version v1 --kind CronJob --conversion
