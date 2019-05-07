package main

import (
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {

}

type rot13Reader struct {
	r io.Reader
}

func Read(r rot13Reader) (int, error) {

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
