package char

import (
	"fmt"
)

var (
	semicolon = byte(';')
	clear     = []byte{27, '[', '2', 'J'}
	home      = []byte{27, '[', '1', semicolon, '1', 'H'}
	escFg     = []byte{27, '[', '3', '8', semicolon, '2', semicolon}
	escBg     = []byte{27, '[', '4', '8', semicolon, '2', semicolon}
)

func locate(pos Position) []byte {
	s := fmt.Sprintf("\x1b[%d;%dH", pos.Y+1, pos.X+1)
	return []byte(s)
}
