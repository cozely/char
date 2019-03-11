package textmode

type Position struct {
	X, Y int
}

func Pos(x, y int) Position {
	return Position{x, y}
}

func (p Position) Plus(other Position) Position {
	return Position{p.X + other.X, p.Y + other.Y}
}

func (p Position) Minus(other Position) Position {
	return Position{p.X - other.X, p.Y - other.Y}
}

func (p Position) In(g Grid) bool {
	return p.X >= g.Min.X && p.X < g.Max.X &&
		p.Y >= g.Min.Y && p.Y < g.Max.Y
}
