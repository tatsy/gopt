package core

import (
	"math"
)

func CoordinateSystem(w *Vector3d) (*Vector3d, *Vector3d) {
	var u, v *Vector3d
	if math.Abs(w.X) > 0.1 {
		u = NewVector3d(0.0, 1.0, 0.0).Cross(w).Normalized()
	} else {
		u = NewVector3d(1.0, 0.0, 0.0).Cross(w).Normalized()
	}
	v = w.Cross(u).Normalized()
	return u, v
}

func SampleCosineHemisphere(u *Vector2d) *Vector3d {
	xy := SampleConcentricDisk(u)
	z := math.Sqrt(math.Max(0.0, (1.0 - (xy.X * xy.X) - (xy.Y * xy.Y))))
	return NewVector3d(xy.X, xy.Y, z)
}

func SampleConcentricDisk(u *Vector2d) *Vector2d {
	uOffset := u.Scale(2.0).Subtract(NewVector2d(1.0, 1.0))
	if uOffset.X == 0.0 && uOffset.Y == 0.0 {
		return NewVector2d(0.0, 0.0)
	}

	var theta, r Float
	if math.Abs(uOffset.X) > math.Abs(uOffset.Y) {
		r = uOffset.X
		theta = (math.Pi * uOffset.Y) / (uOffset.X * 4.0)
	} else {
		r = uOffset.Y
		theta = (math.Pi * 0.5) - (math.Pi*uOffset.X)/(uOffset.Y*4.0)
	}
	return NewVector2d(r*math.Cos(theta), r*math.Sin(theta))
}
