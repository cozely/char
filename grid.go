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

func ScreenGrid() Grid {
	return Grid{Min: Position{0, 0}, Max: Position{width, height}}
}

func (g Grid) Write(pos Position, fg, bg Color, st Style, text []byte) (next Position, err error) {
	cur := g.Min.Plus(pos)

	if !cur.In(g) {
		return pos, errors.New("Put: position out of grid")
	}

	s := norm.NFC.Bytes(text) //TODO: shift responsability to the caller
	for i := 0; i < len(s); {
		d := norm.NFC.NextBoundary(s[i:], true)
		if cur.X >= g.Max.X {
			err = errors.New("Put: line truncated")
			i += d
			continue
		}
		j := cur.X + cur.Y*width
		back.characters[j] = append(back.characters[j][:0], s[i:i+d]...)
		back.foreground[j] = fg
		back.background[j] = bg
		back.style[j] = st
		cur.X++
		i += d
	}

	return cur.Minus(g.Min), err
}
