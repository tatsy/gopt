package core

// Vector2d represents a 2D vector.
type Vector2d struct {
	X Float // x coordinate
	Y Float // y coordinate
}

// NewVector2d returns the pointer to a new Vector2d with specified coordinates.
func NewVector2d(x, y Float) *Vector2d {
	v := new(Vector2d)
	v.X = x
	v.Y = y
	return v
}

// NthElement returns the n-th element of the vector.
func (p *Vector2d) NthElement(i int) Float {
	switch i {
	case 0:
		return p.X
	case 1:
		return p.Y
	default:
		panic("Element index out of range!")
	}
}
