package main

import (
	"net"

	log "github.com/sirupsen/logrus"
)

const port = ":8000"

func main() {
	udpaddr, err := net.ResolveUDPAddr("udp4", port)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpaddr)
	checkError(err)

	for {
		HandleConnection(conn)
	}
}

func HandleConnection(conn *net.UDPConn) {
	var req [512]byte
	n, addr, _ := conn.ReadFromUDP(req[:])

	msg := string(req[0:n])

	log.Info("Received: ", msg)

	// for i := 0; i < 10; i++ {
	// time.Sleep(100 * time.Millisecond)
	_, err := conn.WriteToUDP([]byte(msg), addr)

	checkError(err)
	// }
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
