Instructions
------------

## Pre-requisites

```bash
# set GOPATH & PATH
export GOPATH=~/<GO_DIR_NAME>
export PATH=$PATH:$GOPATH/bin
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
