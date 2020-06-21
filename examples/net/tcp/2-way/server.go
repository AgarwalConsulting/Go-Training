package main

import (
	"net"

	log "github.com/sirupsen/logrus"
)

const servicePort = ":9000"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", servicePort)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		var req [512]byte

		_, err := conn.Read(req[:])

		if err != nil {
			log.Error("Unable to read: ", err)
			break
		}

		log.Info("Request: ", string(req[:]))

		_, err = conn.Write([]byte("pong"))

		if err != nil {
			log.Error("Unable to write: ", err)
			continue
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
