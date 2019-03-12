package main

import (
	"fmt"

	"github.com/cozely/char"
	"golang.org/x/sys/unix"
	"golang.org/x/text/unicode/norm"
)

var (
	white = char.Col(255, 255, 255)
	black = char.Col(0, 0, 0)
)

func main() {
	err := char.Setup()
	if err != nil {
		panic(err)
	}
	defer char.Cleanup()

	scr := char.ScreenGrid()
	grid := scr
	grid.Min.X += 1
	grid.Min.Y += 1
	grid.Max.X -= 1
	grid.Max.Y -= 1

	cur := char.Pos(1, 1)

	var k = []byte{0}
	for {
		s := norm.NFC.Bytes([]byte(fmt.Sprintf("[\u0f3a  %d:é\u032a\u0361:%d \u0f3b ]", cur.X, cur.Y)))

		count := 0
		for i := 0; i < len(s); {
			d := norm.NFC.NextBoundary(s[i:], true)
			count++
			i += d
		}
		_, err = grid.Put(cur, white, black, char.Plain, []byte(s))
		if err != nil {
			scr.Put(char.Pos(0, 0), white, black, char.Plain, []byte(err.Error()))
		} else {
			scr.Put(char.Pos(0, 0), white, black, char.Plain, []byte("----------------------------------------------------------------"))
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
		c := cur
		for i := 0; i < count; i++ {
			c, _ = grid.Put(c, white, black, char.Plain, []byte(" "))
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
