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
	style := textmode.Style{}

	var k = []byte{0}
	for {
		s := fmt.Sprintf("[%02d:%02d]", cur.X, cur.Y)
		_, err = grid.Put(cur, style, []byte(s))
		if err != nil {
			scr.Put(textmode.Pos(0, 0), style, []byte(err.Error()))
		} else {
			scr.Put(textmode.Pos(0, 0), style, []byte("----------------------------------------------------------------"))
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
