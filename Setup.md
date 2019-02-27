Instructions
------------

## Pre-requisites

```bash
# set GOPATH & PATH
export GOPATH=~/<GO_DIR_NAME>
export PATH=$PATH:$GOPATH/bin
```

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

### For a new Project
```bash
# For a new project
dep ensure
```
