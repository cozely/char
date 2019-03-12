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
	return Grid{Min: Position{0, 0}, Max: Position{screen.width, screen.height}}
}

func (g Grid) Put(pos Position, fg, bg Color, st Style, text []byte) (next Position, err error) {
	cur := g.Min.Plus(pos)

	if !cur.In(g) {
		return pos, errors.New("Put: position out of grid")
	}

	count := 0
	s := norm.NFC.Bytes(text) //TODO: shift responsability to the caller
	for i := 0; i < len(s); {
		d := norm.NFC.NextBoundary(s[i:], true)
		count++
		i += d
	}

	if cur.X+count >= g.Max.X {
		if cur.X+count > g.Max.X {
			err = errors.New("Put: line truncated")
			count = g.Max.X - cur.X
		}
	}

	if cur.X < screen.dirty[cur.Y].min {
		screen.dirty[cur.Y].min = cur.X
	}
	if cur.X+count > screen.dirty[cur.Y].max {
		screen.dirty[cur.Y].max = cur.X + count
	}

	var (
		first  int // first byte to be replaced in current line
		trail  int // first byte to be kept in current line
		length int // new length of the slice
	)
	line := &screen.characters[cur.Y]
	for i := 0; i < cur.X; i++ {
		d := norm.NFC.NextBoundary((*line)[first:], true)
		first += d
	}
	trail = first
	for i := 0; i < count; i++ {
		d := norm.NFC.NextBoundary((*line)[trail:], true)
		trail += d
	}
	length = first + len(s) + len(*line) - trail

	if length > cap(*line) {
		l := make([]byte, 0, length+8)
		l = append(l, (*line)[:first]...)
		l = append(l, s...)
		l = append(l, (*line)[trail:]...)
		*line = l
		cur.X += count
		return cur.Minus(g.Min), err
	}

	copy((*line)[first+len(s):first+len(s)+len(*line)-trail], (*line)[trail:])
	*line = (*line)[:length]
	copy((*line)[first:first+len(s)], s)

	cur.X += count
	return cur.Minus(g.Min), err
}
