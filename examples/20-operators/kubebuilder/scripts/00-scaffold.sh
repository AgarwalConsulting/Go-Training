#!/usr/bin/env bash

mkdir project
cd project
# we'll use a domain of tutorial.kubebuilder.io,
# so all API groups will be <group>.tutorial.kubebuilder.io.
kubebuilder init --domain k8s.algogrit.com --repo k8s.algogrit.com/project

# macOS ARM64: https://github.com/kubernetes-sigs/kubebuilder/issues/1932
# kubebuilder init --plugins="kustomize/v2-alpha,base.go.kubebuilder.io/v3"
