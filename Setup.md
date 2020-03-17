# Setup

## Git & Goenv

`goenv` is a go version manager. It allows you to install and work with different versions of Go on the same machine.

```bash
sudo apt-get install git

git clone https://github.com/syndbg/goenv.git ~/.goenv
```

## Pre-requisites

Add to your shell initialization scripts (`~/.bashrc` or `~/.bash_profile` or `~/.zshrc`) the following...

```bash
# set GOPATH & PATH
export GOENV_ROOT="$HOME/.goenv"
export GOPATH=$GOENV_ROOT
export PATH="$GOPATH/bin:$PATH"
export GOENV_DISABLE_GOROOT=1
export GOENV_DISABLE_GOPATH=1
eval "$(goenv init -)"
```

**Important:** Ensure `GOROOT` [isn't set][1]!<sup>[1]</sup>

Follow instructions in: https://github.com/Chennai-Golang/101-workshop/blob/master/README.md

[1]: https://dave.cheney.net/2013/06/14/you-dont-need-to-set-goroot-really

## Working with goenv

### Installing a go version

To list available versions...

```bash
goenv install --list
```

Pick a version and install... Preferably a newer version!

```bash
goenv install {version-number}
```

### Setting a go version globally

```bash
goenv global {version-number}
```

### View all the installed versions

```bash
goenv versions
```

## Test your go installation

```bash
go version
```

### Machine Requirements

- Minimum Specs
  - 8 GB RAM
  - 4-core i5 or higher
- Any OS (MacOS, Windows or Linux)
