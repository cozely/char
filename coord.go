package textmode

type Coord struct {
	X, Y int
}

func Pos(x, y int) Coord {
	return Coord{x, y}
}

func (p Coord) Plus(other Coord) Coord {
	return Coord{p.X + other.X, p.Y + other.Y}
}

func (p Coord) Minus(other Coord) Coord {
	return Coord{p.X - other.X, p.Y - other.Y}
}
