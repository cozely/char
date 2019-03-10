package textmode

import (
	"errors"
	"fmt"
)

type Color uint8

var colorCount = 0

var (
	palette = [256]struct {
		red, green, blue uint8
	}{}
	foreground = [256][]byte{}
	background = [256][]byte{}
)

func NewColor(red, green, blue uint8) (Color, error) {
	if colorCount >= 256 {
		return 0, errors.New("textmode.NewColor: too many colors")
	}

	c := colorCount
	colorCount++
	palette[c].red = red
	palette[c].green = green
	palette[c].blue = blue
	s := fmt.Sprintf(";2;%d;%d;%dm", red, green, blue)
	foreground[c] = []byte("\x1b[38" + s)
	background[c] = []byte("\x1b[48" + s)
	return 0, nil
}
