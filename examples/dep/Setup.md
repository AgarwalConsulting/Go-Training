# `dep` - Go's Package Manager

## Setting up dep

```bash
go get -u github.com/golang/dep/cmd/dep
which dep
# > /Users/gaurav/.goenv/bin/dep
```

## Working with dep

```bash
# Somewhere inside $GOPATH/src
dep init
dep ensure -add github.com/gorilla/mux
```

## For a freshly cloned Project
```bash
# For a freshly cloned project
dep ensure
```
