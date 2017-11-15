package core

import (
	"math"
	"math/rand"
	"testing"
)

func NewRandomVector() *Vector3d {
	v := new(Vector3d)
	v.X = rand.Float64()
	v.Y = rand.Float64()
	v.Z = rand.Float64()
	return v
}

func Equals(u *Vector3d, v *Vector3d) bool {
	return u.X == v.X && u.Y == v.Y && u.Z == v.Z
}

func AlmostEquals(u *Vector3d, v *Vector3d) bool {
	eps := 1.0e-12
	return math.Abs(u.X-v.X) < eps && math.Abs(u.Y-v.Y) < eps && math.Abs(u.Z-v.Z) < eps
}

func TestVector3dInistance(t *testing.T) {
	actual := &Vector3d{}
	expected := NewVector3d(0.0, 0.0, 0.0)
	if !Equals(actual, expected) {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestVector3dConstructWithString(t *testing.T) {
	testCases := map[string][]Float{
		"(1, 2, 3)":    {1.0, 2.0, 3.0},
		"(0, 0, 0)":    {0.0, 0.0, 0.0},
		"(-1, -2, -3)": {-1.0, -2.0, -3.0},
		"(-1, -2, A)":  {},
	}

	for s, vec := range testCases {
		if len(vec) != 0 {
			v := NewVector3dWithString(s)
			if v.X != vec[0] || v.Y != vec[1] || v.Z != vec[2] {
				t.Errorf("%s converted to %v", s, v)
			}
		} else {
			defer func() {
				if p := recover(); p != nil {
				}
			}()
			NewVector3dWithString(s)
			t.Errorf("String \"%s\" must not be parsed.", s)
		}
	}
}

func TestVector3dAdd(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		v := NewRandomVector()
		actual := u.Add(v)
		expected := NewVector3d(u.X+v.X, u.Y+v.Y, u.Z+v.Z)
		if !Equals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dNegate(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		actual := u.Negate()
		expected := NewVector3d(-u.X, -u.Y, -u.Z)
		if !Equals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dSubtract(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		v := NewRandomVector()
		actual := u.Subtract(v)
		expected := NewVector3d(u.X-v.X, u.Y-v.Y, u.Z-v.Z)
		if !Equals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dScale(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		s := rand.Float64()
		actual := u.Scale(s)
		expected := NewVector3d(u.X*s, u.Y*s, u.Z*s)
		if !Equals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dDivide(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		s := rand.Float64()
		actual := u.Divide(s)
		expected := NewVector3d(u.X/s, u.Y/s, u.Z/s)
		if !AlmostEquals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dZeroDivide(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	u := NewRandomVector()
	s := 0.0
	u.Divide(s)
	t.Errorf("Zero division did not panic: %v", u)
}

func TestVector3dDot(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		v := NewRandomVector()
		actual := u.Dot(v)
		expected := u.X*v.X + u.Y*v.Y + u.Z*v.Z
		if actual != expected {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dCross(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		v := NewRandomVector()
		actual := u.Cross(v)
		expected := NewVector3d(
			u.Y*v.Z-u.Z*v.Y,
			u.Z*v.X-u.X*v.Z,
			u.X*v.Y-u.Y*v.X,
		)
		if !AlmostEquals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dLength(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		actual := u.Length()
		expected := math.Sqrt(u.X*u.X + u.Y*u.Y + u.Z*u.Z)
		if actual != expected {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dLengthSquared(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		actual := u.LengthSquared()
		expected := u.X*u.X + u.Y*u.Y + u.Z*u.Z
		if actual != expected {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dNormalized(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		actual := u.Normalized()

		l := math.Sqrt(u.X*u.X + u.Y*u.Y + u.Z*u.Z)
		expected := NewVector3d(
			u.X/l,
			u.Y/l,
			u.Z/l,
		)
		if !AlmostEquals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dZeroNormalized(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	u := Vector3d{0.0, 0.0, 0.0}
	u.Normalized()
	t.Error("Normalizing zero vector did not panic")
}

func TestVector3dMinimum(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		v := NewRandomVector()
		actual := u.Minimum(v)
		expected := NewVector3d(
			math.Min(u.X, v.X),
			math.Min(u.Y, v.Y),
			math.Min(u.Z, v.Z),
		)
		if !Equals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dMaximum(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		v := NewRandomVector()
		actual := u.Maximum(v)
		expected := NewVector3d(
			math.Max(u.X, v.X),
			math.Max(u.Y, v.Y),
			math.Max(u.Z, v.Z),
		)
		if !Equals(actual, expected) {
			t.Errorf("%v expected, but %v detected", expected, actual)
		}
	}
}

func TestVector3dNthElement(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		u := NewRandomVector()
		if u.X != u.NthElement(0) || u.Y != u.NthElement(1) || u.Z != u.NthElement(2) {
			t.Errorf("Element access failed: %v", u)
		}
	}
}

func TestVector3dNthElementOutOfBounds(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		defer func() {
			if r := recover(); r != nil {
			}
		}()

		u := NewRandomVector()
		x := 0
		for x >= 0 && x <= 2 {
			x = rand.Int()
		}

		u.NthElement(x)
		t.Errorf("Out of bounds element access did not panic")
	}
}
