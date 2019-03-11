package char

import (
	"fmt"
)

type Style byte

const (
	Plain       = 0
	Bold  Style = 1 << iota
	Italic
	Underline
	Reverse
	Dim
)

type Color struct {
	R, G, B uint8
}

func Col(red, green, blue uint8) Color {
	return Color{red, green, blue}
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
