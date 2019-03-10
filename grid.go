package textmode

import "golang.org/x/text/unicode/norm"

type Coord struct {
	X, Y int
}

type Grid struct {
	Min, Max   Coord
	Cursor     Coord
	Foreground Color
	Background Color
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
	x, y := g.Min.X+g.Cursor.X, g.Min.Y+g.Cursor.Y

	s := norm.NFC.Bytes(p)
	for i := 0; i < len(s); {
		d := norm.NFC.NextBoundary(s[i:], true)
		cel := &back[y][x]
		cel.glyph = cel.glyph[:0]
		cel.glyph = append(cel.glyph, s[i:i+d]...)
		n++
		x++
		i += d
	}

	g.Cursor.X = x - g.Min.X
	g.Cursor.Y = y - g.Min.Y

	return n, err
}
