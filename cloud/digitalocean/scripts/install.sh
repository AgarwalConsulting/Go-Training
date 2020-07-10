#!/usr/bin/env bash

# set -e

snap install docker doctl tree
snap install kubectl go --classic

mkdir -p .config .kube
sudo snap connect doctl:kube-config
doctl auth init
doctl kubernetes cluster kubeconfig save training-cluster

git clone https://github.com/AgarwalConsulting/Go-Training.git
