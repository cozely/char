package char

import (
	"golang.org/x/sys/unix"
	"golang.org/x/text/unicode/norm"
)

var screen struct {
	width, height int
	characters    [][]byte
	foreground    []Color
	background    []Color
	style         []Style
	dirty         []struct {
		col      int // first dirty character
		min, max int // dirty span in bytes
	}
}

var output []byte

func resize() error {
	var err error

	width, height, err := getSize(unix.Stdout)
	if err != nil {
		return err
	}

	output = make([]byte, width*height+16*height)

	unix.Write(unix.Stdout, escClear)
	unix.Write(unix.Stdout, escHome)
	screen.width, screen.height = width, height
	screen.characters = make([][]byte, height, height)
	screen.foreground = make([]Color, width*height, width*height)
	screen.background = make([]Color, width*height, width*height)
	screen.style = make([]Style, width*height, width*height)
	screen.dirty = make([]struct{ col, min, max int }, height, height)

	for i := range screen.characters {
		screen.characters[i] = make([]byte, width)
		for j := range screen.characters[i] {
			screen.characters[i][j] = byte(' ')
		}
	}
	for i := range screen.foreground {
		screen.foreground[i] = Color{255, 255, 255}
		screen.background[i] = Color{0, 0, 0}
		screen.style[i] = Plain
	}
	for i := range screen.dirty {
		screen.dirty[i].min = 0
		screen.dirty[i].max = 0
	}

	return nil
}

func screenChar(pos Position) []byte {
	n := 0
	for i := 0; i < len(screen.characters[pos.Y]); {
		d := norm.NFC.NextBoundary(screen.characters[pos.Y][i:], true)
		if n == pos.X {
			return screen.characters[pos.Y][i : i+d]
		}
		n++
		i += d
	}
	return nil
}

func screenFg(pos Position) Color {
	return screen.foreground[pos.X+pos.Y*screen.width]
}

func screenBg(pos Position) Color {
	return screen.background[pos.X+pos.Y*screen.width]
}

func screenStyle(pos Position) Style {
	return screen.style[pos.X+pos.Y*screen.width]
}

func Flush() error {
	output = output[:0]
	for line := 0; line < screen.height; line++ {
		dirty := screen.dirty[line]
		if dirty.min == dirty.max {
			continue
		}
		var minbyte, maxbyte int
		for i := 0; i < dirty.min; i++ {
			d := norm.NFC.NextBoundary(screen.characters[line][minbyte:], true)
			minbyte += d
		}
		maxbyte = minbyte
		for i := dirty.min; i < dirty.max; i++ {
			d := norm.NFC.NextBoundary(screen.characters[line][maxbyte:], true)
			maxbyte += d
		}
		output = append(output, locate(Position{dirty.min, line})...)
		output = append(output, screen.characters[line][minbyte:maxbyte]...)
	}
	if len(output) > 0 {
		output = append(output, locate(cursor)...)
		unix.Write(unix.Stdout, output)
	}
	return nil
}
