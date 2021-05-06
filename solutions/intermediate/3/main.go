package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func getRotDictionary() map[byte]byte {
	result := make(map[byte]byte)
	startUpperChar := byte('A')
	startLowerChar := byte('a')

	var i, shiftedChar byte

	for i = 0; i < 26; i++ {
		shiftedChar = (i + 13) % 26
		result[startUpperChar+shiftedChar] = startUpperChar + i
		result[startLowerChar+shiftedChar] = startLowerChar + i
	}

	return result
}

var rotDictionary = getRotDictionary()

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	for idx, el := range p {
		shifted, ok := rotDictionary[el]
		if ok {
			p[idx] = shifted
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
