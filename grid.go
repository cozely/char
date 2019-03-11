package char

import (
	"errors"

	"golang.org/x/text/unicode/norm"
)

type Grid struct {
	Min, Max Position
}

func (g *Grid) Size() Position {
	return g.Max.Minus(g.Min)
}

var screen Grid

func Screen() Grid {
	return screen
}

func (g Grid) Put(pos Position, fg, bg Color, st Style, text []byte) (next Position, err error) {
	cur := g.Min.Plus(pos)

	if !cur.In(g) {
		return pos, errors.New("Put: position out of grid")
	}

	s := norm.NFC.Bytes(text)
	for i := 0; i < len(s); {
		d := norm.NFC.NextBoundary(s[i:], true)

		if !cur.In(g) {
			if err == nil {
				err = errors.New("Put: line truncated")
			}
			break
		}

		cel := &back[cur.Y][cur.X]
		cel.glyph = cel.glyph[:0]
		cel.glyph = append(cel.glyph, s[i:i+d]...)
		cur.X++
		i += d
	}

	return cur.Minus(g.Min), err
}
