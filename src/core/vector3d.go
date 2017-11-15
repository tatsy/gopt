package core

import (
	"fmt"
	"math"
)

// Vector3d is a 3D vector.
type Vector3d struct {
	X, Y, Z Float
}

// NewVector3d creates a new vector with specified coordinates.
func NewVector3d(x, y, z Float) *Vector3d {
	ret := new(Vector3d)
	ret.X = x
	ret.Y = y
	ret.Z = z
	return ret
}

// NewVector3dWithString parses a string, and returns a parsed vector.
func NewVector3dWithString(s string) *Vector3d {
	var x, y, z Float
	n, _ := fmt.Sscanf(s, "(%f, %f, %f)", &x, &y, &z)
	if n != 3 {
		panic(fmt.Sprintf("Failed to parse Vector3d: %s", s))
	}
	return NewVector3d(x, y, z)
}

func (v1 *Vector3d) Add(v2 *Vector3d) *Vector3d {
	ret := new(Vector3d)
	ret.X = v1.X + v2.X
	ret.Y = v1.Y + v2.Y
	ret.Z = v1.Z + v2.Z
	return ret
}

func (v *Vector3d) Negate() *Vector3d {
	ret := new(Vector3d)
	ret.X = -v.X
	ret.Y = -v.Y
	ret.Z = -v.Z
	return ret
}

func (v1 *Vector3d) Subtract(v2 *Vector3d) *Vector3d {
	return v1.Add(v2.Negate())
}

func (v1 *Vector3d) Scale(s Float) *Vector3d {
	ret := new(Vector3d)
	ret.X = v1.X * s
	ret.Y = v1.Y * s
	ret.Z = v1.Z * s
	return ret
}

func (v *Vector3d) Divide(s Float) *Vector3d {
	if s == 0.0 {
		panic("Zero division!")
	}
	return v.Scale(1.0 / s)
}

func (v *Vector3d) Abs() *Vector3d {
	ret := &Vector3d{}
	ret.X = math.Abs(v.X)
	ret.Y = math.Abs(v.Y)
	ret.Z = math.Abs(v.Z)
	return ret
}

func (v1 *Vector3d) Dot(v2 *Vector3d) Float {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 *Vector3d) Cross(v2 *Vector3d) *Vector3d {
	ret := &Vector3d{}
	ret.X = v1.Y*v2.Z - v2.Y*v1.Z
	ret.Y = v1.Z*v2.X - v2.Z*v1.X
	ret.Z = v1.X*v2.Y - v2.X*v1.Y
	return ret
}

func (v1 *Vector3d) Length() Float {
	return Float(math.Sqrt(float64(v1.LengthSquared())))
}

func (v *Vector3d) LengthSquared() Float {
	return v.Dot(v)
}

func (v *Vector3d) Normalized() *Vector3d {
	ret := &Vector3d{}
	ret = v.Divide(v.Length())
	return ret
}

func (v1 *Vector3d) Minimum(v2 *Vector3d) *Vector3d {
	ret := &Vector3d{}
	ret.X = math.Min(v1.X, v2.X)
	ret.Y = math.Min(v1.Y, v2.Y)
	ret.Z = math.Min(v1.Z, v2.Z)
	return ret
}

func (v1 *Vector3d) Maximum(v2 *Vector3d) *Vector3d {
	ret := &Vector3d{}
	ret.X = math.Max(v1.X, v2.X)
	ret.Y = math.Max(v1.Y, v2.Y)
	ret.Z = math.Max(v1.Z, v2.Z)
	return ret
}

func (v *Vector3d) NthElement(i int) Float {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	}
	panic("Element index out of range!")
}

func (v1 *Vector3d) Equals(v2 *Vector3d) bool {
	return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

func (v Vector3d) String() string {
	return fmt.Sprintf("(%.5f, %.5f, %.5f)", v.X, v.Y, v.Z)
}
