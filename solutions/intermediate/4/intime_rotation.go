package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func performRotation(el byte) byte {
	startLowerChar := byte('a')
	endLowerChar := byte('z')

	startUpperChar := byte('A')
	endUpperChar := byte('Z')

	if (startLowerChar <= el) && (el <= endLowerChar) {
		alphabetIdx := ((el - startLowerChar) + 13) % 26
		return startLowerChar + alphabetIdx
	} else if (startUpperChar <= el) && (el <= endUpperChar) {
		alphabetIdx := ((el - startUpperChar) + 13) % 26
		return startUpperChar + alphabetIdx
	} else {
		return el
	}
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for idx, el := range p {
		p[idx] = performRotation(el)
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
