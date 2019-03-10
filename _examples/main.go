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
	grid := scr
	grid.Min.X += 1
	grid.Min.Y += 1
	grid.Max.X -= 1
	grid.Max.Y -= 1
	cur := textmode.Pos(1, 1)

	var k = []byte{0}
	for {
		grid.Cursor = cur
		s := fmt.Sprintf("[%02d:%02d]", cur.X, cur.Y)
		_, err = grid.Write([]byte(s))
		scr.Cursor = textmode.Pos(0, 0)
		if err != nil {
			scr.Write([]byte(err.Error()))
		} else {
			scr.Write([]byte("----------------------------------------------------------------"))
		}
		textmode.Flush()
		//print("GLOP")

		_, err := unix.Read(unix.Stdin, k)
		if err != nil {
			break
		}
		if k[0] == byte('q') || k[0] == 0x11 || k[0] == 0x03 {
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
}
