#!/usr/bin/env bash

cd project
make manifests

# Kustomize v3.8.7 doesn't support arm64
export KUSTOMIZE_INSTALL_SCRIPT=https://gist.githubusercontent.com/algogrit/c3ee63051fcbc685b21d4036fde5ecfb/raw/95025221cbed0bedfbb738280425b4d8ad15db5e/install_kustomize.sh

make install

# export cert_dir=/tmp/k8s-webhook-server/serving-certs/
export cert_dir=/var/folders/4r/tsz32kb54hx8t2m837fndm8r0000gn/T/k8s-webhook-server/serving-certs/

mkdir -p $cert_dir

# https://book.kubebuilder.io/cronjob-tutorial/cert-manager.html
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out $cert_dir/tls.crt \
  -keyout $cert_dir/tls.key \
  -subj "/C=IN/ST=TN/L=Chennai/O=Agarwal Consulting.io/OU=algogrit/CN=localhost"

make run ENABLE_WEBHOOKS=false
