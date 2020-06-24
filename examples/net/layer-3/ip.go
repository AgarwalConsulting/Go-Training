// For raw sockets: https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/raw_sockets_and_the_type_ipconn.html
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("127.0.0.1")

	if ip != nil {
		fmt.Println(ip)
		fmt.Println("Mask", ip.DefaultMask())
	}

	googleIP, _ := net.ResolveIPAddr("ip6", "google.com")

	fmt.Println("IPv6 Google IP: ", googleIP)

	googleIP, _ = net.ResolveIPAddr("ip", "google.com")

	fmt.Println("Google IP: ", googleIP)

	cname, addrs := net.LookupHost("google.com")

	fmt.Println("Host lookup: ", cname, addrs)

	cnames, _ := net.LookupCNAME("algogrit.com")

	fmt.Println("CNAME: ", cnames)
}
