package textmode

import (
	"errors"

	"golang.org/x/sys/unix"
)

var (
	width, height int
	buffer        []byte
)

func resizeBuffer() error {
	var err error

	width, height, err = getSize(unix.Stdout)
	if err != nil {
		return err
	}

	buffer = make([]byte, width*height, width*height)
	for i := range buffer {
		buffer[i] = ' '
	}

	return nil
}

var wbuf []byte

func Put(x, y int, p []byte) (n int, err error) {
	if y < 0 || y >= height || x < 0 || x >= width {
		return 0, errors.New("textmode.Write: out of screen")
	}

	wbuf = wbuf[:0]
	wbuf = append(wbuf, locate(x, y)...)
	//TODO: unicode support
	for i := range p {
		switch {
		case x >= width:
			if err == nil {
				err = errors.New("textmode.Write: line truncated")
			}
		case y >= height:
			if err == nil {
				err = errors.New("textmode.Write: line out of screen")
			}
			return n, err
		case p[i] == buffer[x+y*width]:
		default:
			wbuf = append(wbuf, p[i])
			n++
			x++
		}
	}

	if n > 0 {
		unix.Write(unix.Stdout, wbuf)
	}

	return n, err
}
