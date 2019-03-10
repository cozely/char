package main

import (
	"fmt"

	"github.com/cozely/textmode"
	"golang.org/x/sys/unix"
)

func main() {
	err := textmode.Setup()
	if err != nil {
		panic(err)
	}
	defer textmode.Cleanup()

	scr := textmode.Screen()
	cur := textmode.Coord{X: 2, Y: 2}

	var k = []byte{0}
	for {
		scr.Cursor = cur
		s := fmt.Sprintf("[%02d:%02d]", cur.X, cur.Y)
		_, err = scr.Write([]byte(s))
		if err != nil {
			panic(err)
		}
		textmode.Flush()
		//print("GLOP")

		_, err := unix.Read(unix.Stdin, k)
		if err != nil {
			break
		}
		if k[0] == byte('Q') || k[0] == 0x11 || k[0] == 0x03 {
			break
		}
		switch k[0] {
		case 'h':
			cur.X--
		case 'j':
			cur.Y++
		case 'k':
			cur.Y--
		case 'l':
			cur.X++
		}
	}

	unix.Write(unix.Stdout, []byte("Bye!u\u0312u \u0312 u\u0323u\n\r"))
}
