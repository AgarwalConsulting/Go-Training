package main

import (
	"flag"
	"log"
)

func main() {
	var ip = flag.Int("myFlag", 1234, "help message for myFlag")

	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Option myFlag: %T, %#v, %d\n", ip, ip, *ip)
}
