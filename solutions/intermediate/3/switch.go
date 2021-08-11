package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for idx, el := range p {
		switch {
		case el+13 > 'z':
			p[idx] = el - 13
		case el >= 'a':
			p[idx] = el + 13
		case el+13 > 'Z':
			p[idx] = el - 13
		case el >= 'A':
			p[idx] = el + 13
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
