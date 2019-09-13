Instructions
------------

## Git & Goenv

```bash
sudo apt-get install git

git clone https://github.com/syndbg/goenv.git ~/.goenv
```

## Pre-requisites

Add to your shell initialization scripts (`~/.bashrc` or `~/.bash_profile` or `~/.zshrc`) the following...

```bash
# set GOPATH & PATH
export GOENV_ROOT="$HOME/.goenv"
export GOPATH=~/$GOENV_ROOT
export PATH=$PATH:$GOPATH/bin
export GOENV_DISABLE_GOROOT=1
export GOENV_DISABLE_GOPATH=1
eval "$(goenv init -)"
```

**Important:** Ensure `GOROOT` isn't set!<sup>[1]</sup>

Follow instructions in: https://github.com/AgarwalConsulting/learning-golang/blob/master/README.md

## `dep` - Go's Package Manager

### Setting up dep

```bash
go get -u github.com/golang/dep/cmd/dep
which dep
# > /Users/gaurav/.goenv/bin/dep
```

### Working with dep

```bash
# Somewhere inside $GOPATH/src
dep init
dep ensure -add github.com/gorilla/mux
```

### For a freshly cloned Project
```bash
# For a freshly cloned project
dep ensure
```

[1]: https://dave.cheney.net/2013/06/14/you-dont-need-to-set-goroot-really
