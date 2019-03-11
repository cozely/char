package char

import (
	"reflect"
	"unsafe"
)

type Cell struct {
	character  character
	foreground Color
	background Color
}

/*
func (c Cell) Equals(other Cell) bool {
	if c.foreground != other.foreground || c.background != other.background {
		return false
	}
	for i := range c.glyph {
		if i >= len(other.glyph) {
			return false
		}
		if c.glyph[i] != other.glyph[i] {
			return false
		}
	}

	return true
}
*/

type character string

var (
	multibyte = make(map[string]character)
	ascii     = [128]character{
		"\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd",
		"\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd",
		"\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd",
		"\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd", "\ufffd",
		" ", "!", `"`, "#", "$", "%", "&", "'",
		"(", ")", "*", "+", ",", "-", ".", "/",
		"0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9", ":", ";", "<", "=", ">", "?",
		"@", "A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z", "[", `\`, "]", "^", "_",
		"`", "a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t", "u", "v", "w",
		"x", "y", "z", "{", "|", "}", "~", "\ufffd",
	}
)

// Intern returns the internalized string corresponding to a single character.
// The byte slice must be a valid UTF8 sequence, NFC normalized, and correspond
// to a single character (it may contains combining code points).
func Intern(b []byte) character {
	if b[0] < 128 {
		return ascii[b[0]]
	}
	sl := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	str := reflect.StringHeader{Data: sl.Data, Len: sl.Len}
	s := *(*string)(unsafe.Pointer(&str))
	result, ok := multibyte[s]
	if !ok {
		multibyte[s] = character(s)
		return character(s)
	}
	return result
}
