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

// Add computes the sum of two vectors.
func (v1 *Vector2d) Add(v2 *Vector2d) *Vector2d {
	ret := new(Vector2d)
	ret.X = v1.X + v2.X
	ret.Y = v1.Y + v2.Y
	return ret
}

// Subtract computes the difference of two vectors.
func (v1 *Vector2d) Subtract(v2 *Vector2d) *Vector2d {
	ret := new(Vector2d)
	ret.X = v1.X - v2.X
	ret.Y = v1.Y - v2.Y
	return ret
}

// Scale returns the scaled vector.
func (v *Vector2d) Scale(s Float) *Vector2d {
	ret := new(Vector2d)
	ret.X = v.X * s
	ret.Y = v.Y * s
	return ret
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
