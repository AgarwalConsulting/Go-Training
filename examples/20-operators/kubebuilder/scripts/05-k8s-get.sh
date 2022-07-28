#!/usr/bin/env bash

kubectl get cronjob.batch.k8s.algogrit.com -o yaml
kubectl get jobs -o yaml
kubectl get pods -o yaml
