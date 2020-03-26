package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type ErrorWrapper struct {
	info string
	err  error
}

func (e ErrorWrapper) Unwrap() error {
	return e.err
}

func (e ErrorWrapper) Error() string {
	return e.info + ": " + e.err.Error()
}

func Read(fileName string) (contents string, err error) {
	file, err := os.Open(fileName)

	if err != nil {
		message := fileName + " doesn't exists. Please check the location."
		return "", ErrorWrapper{message, err}
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		message := "scan error!"

		return "", ErrorWrapper{message, err}
	}

	return
}

func main() {
	contents, err := Read(os.Args[0])

	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	fmt.Println(contents)
}
