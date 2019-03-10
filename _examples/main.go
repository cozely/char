package main

import (
	"github.com/cozely/textmode"
	"golang.org/x/sys/unix"
)

func main() {
	err := textmode.Setup()
	if err != nil {
		panic(err)
	}
	defer textmode.Cleanup()

	x, y := 2, 2
	var k = []byte{0}
	for {
		_, err = textmode.Put(x, y, []byte("Hello,u\u0306u"))
		if err != nil {
			panic(err)
		}
		_, err = textmode.Put(x+1, y+1, []byte("World!"))
		if err != nil {
			panic(err)
		}

		_, err := unix.Read(unix.Stdin, k)
		if err != nil {
			break
		}
		if k[0] == byte('Q') || k[0] == 0x11 || k[0] == 0x03 {
			break
		}
		switch k[0] {
		case 'h':
			x--
		case 'j':
			y++
		case 'k':
			y--
		case 'l':
			x++
		}
	}

	unix.Write(unix.Stdout, []byte("Bye!u\u0312u \u0312 u\u0323u\n\r"))
}
