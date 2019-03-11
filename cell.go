package char

type Cell struct {
	glyph      []byte
	foreground Color
	background Color
}

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
