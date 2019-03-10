// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux,!appengine netbsd openbsd

// Package terminal provides support functions for dealing with terminals, as
// commonly found on UNIX systems.
//
// Putting a terminal into raw mode is the most common requirement:
//
// 	oldState, err := terminal.MakeRaw(0)
// 	if err != nil {
// 	        panic(err)
// 	}
// 	defer terminal.Restore(0, oldState)
package textmode

import "golang.org/x/sys/unix"

// state contains the state of a terminal.
type state struct {
	termios unix.Termios
}

// isTerminal returns true if the given file descriptor is a terminal.
func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}

// makeRaw put the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func makeRaw(fd int) (*state, error) {
	termios, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	if err != nil {
		return nil, err
	}

	oldState := state{termios: *termios}

	// This attempts to replicate the behaviour documented for cfmakeraw in
	// the termios(3) manpage.
	termios.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	termios.Oflag &^= unix.OPOST
	termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	termios.Cflag &^= unix.CSIZE | unix.PARENB
	termios.Cflag |= unix.CS8
	termios.Cc[unix.VMIN] = 1
	termios.Cc[unix.VTIME] = 0
	if err := unix.IoctlSetTermios(fd, ioctlWriteTermios, termios); err != nil {
		return nil, err
	}

	return &oldState, nil
}

// getState returns the current state of a terminal which may be useful to
// restore the terminal after a signal.
func getState(fd int) (*state, error) {
	termios, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	if err != nil {
		return nil, err
	}

	return &state{termios: *termios}, nil
}

// restore restores the terminal connected to the given file descriptor to a
// previous state.
func restore(fd int, s *state) error {
	return unix.IoctlSetTermios(fd, ioctlWriteTermios, &s.termios)
}

// getSize returns the dimensions of the given terminal.
func getSize(fd int) (width, height int, err error) {
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	return int(ws.Col), int(ws.Row), nil
}
