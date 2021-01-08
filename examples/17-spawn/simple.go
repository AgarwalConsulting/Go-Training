package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strconv"
	"sync"
)

func main() {
	// cmd := exec.Command("ls", "-lrth", "/tmp/")
	cmds := make(map[*exec.Cmd]io.ReadCloser)

	for i := 0; i < 10; i++ {
		cmd := exec.Command("sleep", strconv.Itoa(i))
		outPipe, _ := cmd.StdoutPipe()

		cmd.Start()

		cmds[cmd] = outPipe
	}

	// outPipe, _ := cmd.StdoutPipe()

	// cmd.Start()

	fmt.Println("Hello, World!")
	fmt.Println("Doing some busy work!")

	var wg sync.WaitGroup
	for cmd, outPipe := range cmds {
		wg.Add(1)
		go func(cmd *exec.Cmd, outPipe io.ReadCloser) {
			defer wg.Done()
			out, _ := ioutil.ReadAll(outPipe)
			cmd.Wait()
			fmt.Println(string(out))
		}(cmd, outPipe)
	}

	wg.Wait()
	// out, _ := cmd.Output()
}
