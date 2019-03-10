// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build solaris

package textmode

import (
	"golang.org/x/sys/unix"
	"io"
	"syscall"
)

// state contains the state of a terminal.
type state struct {
	termios unix.Termios
}

// isTerminal returns true if the given file descriptor is a terminal.
func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermio(fd, unix.TCGETA)
	return err == nil
}

// makeRaw puts the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
// see http://cr.illumos.org/~webrev/andy_js/1060/
func makeRaw(fd int) (*state, error) {
	termios, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return nil, err
	}

	oldState := State{termios: *termios}

	termios.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	termios.Oflag &^= unix.OPOST
	termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	termios.Cflag &^= unix.CSIZE | unix.PARENB
	termios.Cflag |= unix.CS8
	termios.Cc[unix.VMIN] = 1
	termios.Cc[unix.VTIME] = 0

	if err := unix.IoctlSetTermios(fd, unix.TCSETS, termios); err != nil {
		return nil, err
	}

	return &oldState, nil
}

// restore restores the terminal connected to the given file descriptor to a
// previous state.
func restore(fd int, oldState *state) error {
	return unix.IoctlSetTermios(fd, unix.TCSETS, &oldState.termios)
}

// getState returns the current state of a terminal which may be useful to
// restore the terminal after a signal.
func getState(fd int) (*state, error) {
	termios, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return nil, err
	}

	return &state{termios: *termios}, nil
}

// getSize returns the dimensions of the given terminal.
func getSize(fd int) (width, height int, err error) {
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	return int(ws.Col), int(ws.Row), nil
}
