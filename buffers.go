package char

import (
	"golang.org/x/sys/unix"
)

type buffer [][]Cell

func newBuffer(width, height int) buffer {
	b := make([][]Cell, height, height)
	for i := range b {
		b[i] = make([]Cell, width, width)
		for j := range b[i] {
			b[i][j] = Cell{
				glyph:      []byte{byte(' ')},
				foreground: Color{R: 255, G: 255, B: 255},
				background: Color{R: 0, G: 0, B: 0},
			}
		}
	}
	return b
}

func (b buffer) At(p Position) Cell {
	return b[p.Y][p.X]
}

func (b buffer) Set(p Position, c Cell) {
	cel := &b[p.Y][p.X]
	cel.glyph = cel.glyph[:0]
	cel.glyph = append(cel.glyph, c.glyph...)
	cel.foreground = c.foreground
	cel.background = c.background
}

var (
	front  buffer
	back   buffer
	output []byte
)

func resize() error {
	var err error

	width, height, err := getSize(unix.Stdout)
	if err != nil {
		return err
	}

	front = newBuffer(width, height)
	back = newBuffer(width, height)
	output = make([]byte, 2*width*height)

	unix.Write(unix.Stdout, clear)
	screen.Min = Position{0, 0}
	screen.Max = Position{width, height}

	return nil
}

func Flush() error {
	output = output[:0]
	cur := Position{X: -1, Y: -1}
	var pos Position
	for pos.Y = range front {
		for pos.X = range front[pos.Y] {
			cel := back.At(pos)
			if front.At(pos).Equals(cel) {
				continue
			}
			if cur != pos {
				output = append(output, locate(pos)...)
				cur = pos
			}
			output = append(output, cel.glyph...)
			front.Set(pos, cel)
			cur.X++
		}
	}
	if len(output) > 0 {
		output = append(output, locate(cursor)...)
		unix.Write(unix.Stdout, output)
	}
	return nil
}
