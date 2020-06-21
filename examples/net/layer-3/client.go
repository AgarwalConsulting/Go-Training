package main

import (
	"net"

	log "github.com/sirupsen/logrus"
)

const service = "127.0.0.1"

func main() {
	addr, err := net.ResolveIPAddr("ip", service)
	checkError(err)
	conn, err := net.DialIP("ip:udp", nil, addr)
	checkError(err)

	conn.Write([]byte("hello"))

	var resp [512]byte
	n, _ := conn.Read(resp[:])

	log.Info("Received: ", string(resp[0:n]))
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}
}
