package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b <= byte('z') && b >= byte('n') || b <= byte('Z') && b >= byte('N') {
		return b - 13
	} else if b <= byte('m') && b >= byte('a') || b <= byte('M') && b >= byte('A') {
		return b + 13
	}
	return b
}
func (r rot13Reader) Read(b []byte) (int, error) {
	buf := make([]byte, len(b), cap(b))
	bits, err := r.r.Read(buf)
	for i := 0; i < bits; i++ {
		b[i] = rot13(buf[i])
	}
	return bits, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
