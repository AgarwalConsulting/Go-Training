#!/usr/bin/env bash

set -e

# Install the latest version of tools
apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -y aptitude bridge-utils python2.7 python-setuptools apt-transport-https ca-certificates curl wget time software-properties-common python3-pip virtualenv python3-setuptools gnupg gnupg-agent tcpdump ngrep tshark silversearcher-ag-el bpfcc-tools ipython fzf vim dnsutils jq golang-go git libpcap0.8 libpcap0.8-dev tmux bpfcc-tools bpftool

# Install and configure docker
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -y docker-ce docker-ce-cli containerd.io
systemctl enable docker

# Install kind
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.11.1/kind-linux-amd64
chmod +x ./kind
mv ./kind /usr/local/bin/

# Install footloose
curl -Lo footloose https://github.com/weaveworks/footloose/releases/download/0.6.3/footloose-0.6.3-linux-x86_64
chmod +x footloose
mv footloose /usr/local/bin/

# Install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Install kubebuilder
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/

# # Install Operator SDK
# export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac)
# export OS=$(uname | awk '{print tolower($0)}')
