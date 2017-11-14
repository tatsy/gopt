package core

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandVector() *Vector3d {
	v := &Vector3d{}
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

func TestInistance(t *testing.T) {
	actual := &Vector3d{}
	expected := &Vector3d{0.0, 0.0, 0.0}
	if !Equals(actual, expected) {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestAdd(t *testing.T) {
	u := RandVector()
	v := RandVector()
	actual := u.Add(v)
	expected := NewVector3d(u.X+v.X, u.Y+v.Y, u.Z+v.Z)
	if !Equals(actual, expected) {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestNegate(t *testing.T) {
	u := RandVector()
	actual := u.Negate()
	expected := NewVector3d(-u.X, -u.Y, -u.Z)
	if !Equals(actual, expected) {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	u := RandVector()
	v := RandVector()
	actual := u.Subtract(v)
	expected := NewVector3d(u.X-v.X, u.Y-v.Y, u.Z-v.Z)
	if !Equals(actual, expected) {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestScale(t *testing.T) {
	u := RandVector()
	s := rand.Float64()
	actual := u.Scale(s)
	expected := NewVector3d(u.X*s, u.Y*s, u.Z*s)
	if !Equals(actual, expected) {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestDivide(t *testing.T) {
	u := RandVector()
	s := rand.Float64()
	actual := u.Divide(s)
	expected := NewVector3d(u.X/s, u.Y/s, u.Z/s)
	if !AlmostEquals(actual, expected) {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestZeroDivide(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	u := RandVector()
	s := 0.0
	u.Divide(s)
	t.Errorf("Zero division did not panic: %v", u)
}

func TestDot(t *testing.T) {
	u := RandVector()
	v := RandVector()
	actual := u.Dot(v)
	expected := u.X*v.X + u.Y*v.Y + u.Z*v.Z
	if actual != expected {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestCross(t *testing.T) {
	u := RandVector()
	v := RandVector()
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

func TestLength(t *testing.T) {
	u := RandVector()
	actual := u.Length()
	expected := math.Sqrt(u.X*u.X + u.Y*u.Y + u.Z*u.Z)
	if actual != expected {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestLengthSquared(t *testing.T) {
	u := RandVector()
	actual := u.LengthSquared()
	expected := u.X*u.X + u.Y*u.Y + u.Z*u.Z
	if actual != expected {
		t.Errorf("%v expected, but %v detected", expected, actual)
	}
}

func TestNormalized(t *testing.T) {
	u := RandVector()
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

func TestZeroNormalized(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	u := Vector3d{0.0, 0.0, 0.0}
	u.Normalized()
	t.Error("Normalizing zero vector did not panic")
}

func TestMinimum(t *testing.T) {
	u := RandVector()
	v := RandVector()
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

func TestMaximum(t *testing.T) {
	u := RandVector()
	v := RandVector()
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

func TestNthElement(t *testing.T) {
	u := RandVector()
	if u.X != u.NthElement(0) || u.Y != u.NthElement(1) || u.Z != u.NthElement(2) {
		t.Errorf("Element access failed: %v", u)
	}
}

func TestNthElementOutOfBounds(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	u := RandVector()
	x := 0
	for x >= 0 && x <= 2 {
		x = rand.Int()
	}

	u.NthElement(x)
	t.Errorf("Out of bounds element access did not panic")
}
