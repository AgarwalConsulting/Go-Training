package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func substitute(c byte) byte {
	if c >= 'a' && c < 'n' {
		return c + 13
	}
	if c >= 'n' && c <= 'z' {
		return c - 13
	}
	if c >= 'A' && c < 'N' {
		return c + 13
	}
	if c >= 'N' && c <= 'Z' {
		return c - 13
	}

	return c
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)

	for idx, el := range p {
		p[idx] = substitute(el)
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
