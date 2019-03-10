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

func Put(x, y int, p []byte) (n int, err error) {
	if y < 0 || y >= height || x < 0 || x >= width {
		return 0, errors.New("textmode.Write: out of screen")
	}

	for n < len(p) {
		//TODO: unicode support

		// Skip content that is already on screen
		for p[n] == buffer[x+y*width] {
			n++
			if n == len(p) {
				return n, nil
			}
			x++
			if x == width {
				return n, errors.New("textmode.Write: truncated")
			}
		}

		count := 0
		for p[n+count] != buffer[x+count+y*width] {
			count++
			if n+count == len(p) {
				break
			}
			if x+count == width {
				err = locate(x, y)
				if err != nil {
					return n, err //TODO: wrap
				}
				nn, err := unix.Write(unix.Stdout, p[n:n+count])
				if err != nil {
					return n + nn, err //TODO: wrap
				}
				return n + count, errors.New("textmode.Write: truncated")
			}
		}
		err = locate(x, y)
		if err != nil {
			return n, err //TODO: wrap
		}
		nn, err := unix.Write(unix.Stdout, p[n:n+count])
		if err != nil {
			return n + nn, err //TODO: wrap
		}
		n += count
		x += count
	}

	return n, nil
}
