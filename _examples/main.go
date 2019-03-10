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
	msg := "Ok "

	var k = []byte{0}
	for {
		scr.Cursor = cur
		s := fmt.Sprintf("[%02d:%02d:%s]", cur.X, cur.Y, msg)
		_, err = scr.Write([]byte(s))
		if err != nil {
			msg = "Err"
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
