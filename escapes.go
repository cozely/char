package textmode

import (
	"fmt"

	"golang.org/x/sys/unix"
)

var (
	clear = []byte{27, '[', '2', 'J'}
	home  = []byte{27, '[', '1', ';', '1', 'H'}
)

func locate(x, y int) error {
	s := fmt.Sprintf("\x1b[%d;%dH", y+1, x+1)
	_, err := unix.Write(unix.Stdout, []byte(s))
	if err != nil {
		return err
	}
	return nil
}
