# Setup


## Steps

### Step 1

Verify if you have [git](https://git-scm.com/) & [golang](http://golang.org/) installed.

From your shell (`git bash` for windows users), run

```bash
go version
```

Make sure you have `go` version `1.13` or above installed.

### Step 2

Verify the `GOPATH`.

```bash
echo $GOPATH
```

Your `GOPATH` is where you would be storing and managing your project and any third-party source code.

#### Step 2.1

In case, you don't the `GOPATH` set, then set the variable to the full-path to somewhere in your file-system.

### For *Nix

Add to your shell initialization scripts (`~/.bashrc` or `~/.bash_profile` or `~/.zshrc`) the following...

```bash
# set GOPATH & PATH
export GOPATH=~/go # Or some other path in your system
export PATH="$GOPATH/bin:$PATH"
```

P.S. You can check which shell you are using by running `echo $SHELL`.

### For Windows

You would have to Edit your "Environment Variables". Refer: [How to set environment variables](https://docs.oracle.com/en/database/oracle/machine-learning/oml4r/1.5.1/oread/creating-and-modifying-environment-variables-on-windows.html#GUID-DD6F9982-60D5-48F6-8270-A27EC53807D0)

Ensure `%GOPATH%` is set and `%GOPATH%\bin` is added to the PATH.

> Restart your `git bash` and verify the steps 1 and 2 again.

### Step 3

Follow `README.md` to get Go-Training.

```bash
mkdir -p $GOPATH/src/github.com/AgarwalConsulting

git clone https://github.com/AgarwalConsulting/Go-Training.git $GOPATH/src/github.com/AgarwalConsulting/Go-Training

cd $GOPATH/src/github.com/AgarwalConsulting/Go-Training
```

(and)

```bash
go install github.com/AgarwalConsulting/Go-Training@latest
```

### Step 4

Verify your installation

```bash
Go-Training
```

should give you an output of:

```md
Hello World! Result:  92
0x?????????? 0x??????????
```

## Machine Requirements

- Minimum Specs
  - 8 GB RAM
  - 4-core i5 or higher
- Any OS (MacOS, Windows or Linux)
