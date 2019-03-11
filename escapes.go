package char

import (
	"fmt"
)

var (
	escSmcup      = []byte("\x1b[?1049h")
	escRmcup      = []byte("\x1b[?1049l")
	escClear      = []byte("\x1b[2J")
	escHome       = []byte("\x1b[1;1H")
	escShowCursor = []byte("\x1b[?25l")
	escHideCursor = []byte("\x1b[?25h")
	escFg         = []byte("\x1b[38;2;")
	escBg         = []byte("\x1b[48;2;")
)

func locate(pos Position) []byte {
	s := fmt.Sprintf("\x1b[%d;%dH", pos.Y+1, pos.X+1)
	return []byte(s)
}

func cursorShape(shape int) []byte {
	s := fmt.Sprintf("\x1b[%d q", shape)
	// For linux console: "\x1b[?%dc"
	return []byte(s)
}
