package main

import (
	log "github.com/sirupsen/logrus"
)

const port = ":8000"

func main() {
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}
}
