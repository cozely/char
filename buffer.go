package char

import (
	"bytes"

	"golang.org/x/sys/unix"
)

type buffer struct {
	characters [][]byte
	foreground []Color
	background []Color
	style      []Style
}

var (
	width, height int
	front, back   buffer
	output        []byte
)

func newBuffer() buffer {
	b := buffer{
		characters: make([][]byte, width*height, width*height),
		foreground: make([]Color, width*height, width*height),
		background: make([]Color, width*height, width*height),
		style:      make([]Style, width*height, width*height),
	}

	for i := range b.characters {
		b.characters[i] = []byte{byte(' ')}
		b.foreground[i] = Color{255, 255, 255}
		b.background[i] = Color{0, 0, 0}
		b.style[i] = Plain
	}

	return b
}

func resize() error {
	var err error

	width, height, err = getSize(unix.Stdout)
	if err != nil {
		return err
	}

	output = make([]byte, width*height+16*height)
	front = newBuffer()
	back = newBuffer()

	unix.Write(unix.Stdout, escClear)
	unix.Write(unix.Stdout, escHome)

	return nil
}

func Flush() error {
	output = output[:0]
	cursor := Position{-1, -1}
	var pos Position
	for pos.Y = 0; pos.Y < height; pos.Y++ {
		for pos.X = 0; pos.X < width; pos.X++ {
			i := pos.X + pos.Y*width
			fg := back.foreground[i] == front.foreground[i]
			bg := back.background[i] == front.background[i]
			style := back.style[i] != front.style[i]
			character := !bytes.Equal(back.characters[i], front.characters[i])
			if fg || bg || style || character {
				if cursor != pos {
					output = append(output, locate(pos)...)
					cursor = pos
				}
				output = append(output, back.characters[i]...)
				front.characters[i] = append(front.characters[i][:0], back.characters[i]...)
				cursor.X++
			}
		}
	}

	if len(output) > 0 {
		output = append(output, locate(screenCursor)...)
		unix.Write(unix.Stdout, output)
	}
	return nil
}
