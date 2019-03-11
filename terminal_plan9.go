// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
package char

import (
	"fmt"
	"runtime"
)

type state struct{}

// isTerminal returns true if the given file descriptor is a terminal.
func isTerminal(fd int) bool {
	return false
}

// makeRaw put the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func makeRaw(fd int) (*state, error) {
	return nil, fmt.Errorf("terminal: MakeRaw not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}

// getState returns the current state of a terminal which may be useful to
// restore the terminal after a signal.
func getState(fd int) (*state, error) {
	return nil, fmt.Errorf("terminal: GetState not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}

// restore restores the terminal connected to the given file descriptor to a
// previous state.
func restore(fd int, s *state) error {
	return fmt.Errorf("terminal: Restore not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}

// getSize returns the dimensions of the given terminal.
func getSize(fd int) (width, height int, err error) {
	return 0, 0, fmt.Errorf("terminal: GetSize not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}
