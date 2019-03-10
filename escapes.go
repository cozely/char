package textmode

import (
	"fmt"
)

var (
	clear = []byte{27, '[', '2', 'J'}
	home  = []byte{27, '[', '1', ';', '1', 'H'}
)

func locate(x, y int) []byte {
	s := fmt.Sprintf("\x1b[%d;%dH", y+1, x+1)
	return []byte(s)
}
