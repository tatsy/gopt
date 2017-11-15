package bsdf

import (
	"math"

	. "github.com/tatsy/gopt/src/core"
)

type LambertReflection struct {
	re *Color
}

func NewLambertReflection(re *Color) *LambertReflection {
	f := &LambertReflection{}
	f.re = re
	return f
}

func (f *LambertReflection) Eval(wi, wo *Vector3d) *Color {
	return f.re.Scale(1.0 / math.Pi)
}

func (f *LambertReflection) Pdf(wi, wo *Vector3d) Float {
	if SameHemisphere(wi, wo) {
		return AbsCosTheta(wi) / math.Pi
	}
	return 0.0
}

func (f *LambertReflection) Sample(wo *Vector3d, u *Vector2d) (*Color, *Vector3d, Float) {
	phi := 2.0 * math.Pi * u.X
	r2 := u.Y
	r := math.Sqrt(r2)

	x := math.Cos(phi) * r
	y := math.Sin(phi) * r
	z := math.Sqrt(1.0 - r2)
	return f.re, NewVector3d(x, y, z), 1.0
}

func (f *LambertReflection) Type() int {
	return BSDF_DIFFUSE | BSDF_REFLECTION
}
