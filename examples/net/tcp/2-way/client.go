package main

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

const servicePort = ":9000"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", servicePort)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	conn.SetKeepAlive(true)

	for {
		msg := []byte("ping")
		_, err := conn.Write(msg)

		if err != nil {
			log.Error("Unable to write: ", err)
			break
		}

		var resp [512]byte

		_, err = conn.Read(resp[:])

		if err != nil {
			log.Error("Unable to read: ", err)
			continue
		}

		time.Sleep(time.Millisecond * 100)

		log.Info("Response: ", string(resp[:]))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
