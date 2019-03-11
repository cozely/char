package textmode

import (
	"fmt"
)

type Style struct {
	Fg    Color
	Bg    Color
	Attrb uint8
}

type Color struct {
	R, G, B uint8
}

var colors map[Color][]byte

func (c Color) bytes() []byte {
	b, ok := colors[c]
	if !ok {
		b = []byte(fmt.Sprintf("%d;%d;%dm", c.R, c.G, c.B))
		colors[c] = b
	}
	return b
}
