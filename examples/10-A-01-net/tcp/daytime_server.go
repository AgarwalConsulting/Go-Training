// Sample from: https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/tcp_sockets.html
package main

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		var req [512]byte
		conn.Read(req[:])
		log.Info(string(req))

		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Write([]byte("\n"))
		conn.Close() // we're finished with this client
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
