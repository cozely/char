package textmode

import (
	"errors"

	"golang.org/x/text/unicode/norm"
)

type Grid struct {
	Min, Max   Coord
	Cursor     Coord
	Foreground Color
	Background Color
}

func (g *Grid) Contains(pos Coord) bool {
	return pos.X >= g.Min.X && pos.X < g.Max.X &&
		pos.Y >= g.Min.Y && pos.Y < g.Max.Y
}

var screen = Grid{
	Cursor:     Coord{0, 0},
	Foreground: Color{255, 255, 255},
	Background: Color{0, 0, 0},
}

func Screen() Grid {
	return screen
}

func (g *Grid) Write(p []byte) (n int, err error) {
	cur := g.Min.Plus(g.Cursor)

	s := norm.NFC.Bytes(p)
	for i := 0; i < len(s); {
		d := norm.NFC.NextBoundary(s[i:], true)
		if g.Contains(cur) {
			cel := &back[cur.Y][cur.X]
			cel.glyph = cel.glyph[:0]
			cel.glyph = append(cel.glyph, s[i:i+d]...)
			n++
		} else if err == nil {
			err = errors.New("out of grid")
		}
		cur.X++
		i += d
	}

	g.Cursor = cur.Minus(g.Min)

	return n, err
}
