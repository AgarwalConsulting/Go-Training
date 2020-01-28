package main

// // No C code needed.
import "C"

import "log"

var W int

func F() { log.Printf("Hello, from logger plugin %d\n", W) }
