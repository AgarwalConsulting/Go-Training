#!/usr/bin/env bash
# https://sdk.operatorframework.io/docs/installation/

rm -rf myoperator
mkdir myoperator
cd myoperator
operator-sdk init --domain=algogrit.com --repo=github.com/AgarwalConsulting/myplatform operator --skip-go-version-check
