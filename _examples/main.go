package main

import (
	"fmt"

	"github.com/cozely/char"
	"golang.org/x/sys/unix"
)

func main() {
	err := char.Setup()
	if err != nil {
		panic(err)
	}
	defer char.Cleanup()

	scr := char.Screen()
	grid := scr
	grid.Min.X += 1
	grid.Min.Y += 1
	grid.Max.X -= 1
	grid.Max.Y -= 1

	cur := char.Pos(1, 1)
	style := char.Style{}

	var k = []byte{0}
	for {
		s := fmt.Sprintf("[%02d:%02d]", cur.X, cur.Y)
		_, err = grid.Put(cur, style, []byte(s))
		if err != nil {
			scr.Put(char.Pos(0, 0), style, []byte(err.Error()))
		} else {
			scr.Put(char.Pos(0, 0), style, []byte("----------------------------------------------------------------"))
		}
		char.Flush()
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
