#!/usr/bin/env bash

cd myoperator
operator-sdk create api --group=myplatform --version=v1alpha1 --kind=InhouseApp
