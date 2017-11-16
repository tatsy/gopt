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
	if SameHemisphere(wi, wo) {
		return f.re.Scale(1.0 / math.Pi)
	}
	return NewColor(0.0, 0.0, 0.0)
}

func (f *LambertReflection) Pdf(wi, wo *Vector3d) Float {
	if SameHemisphere(wi, wo) {
		return AbsCosTheta(wi) / math.Pi
	}
	return 0.0
}

func (f *LambertReflection) Sample(wo *Vector3d, u *Vector2d) (*Color, *Vector3d, Float, BsdfType) {
	wi := SampleCosineHemisphere(u)
	fr := f.Eval(wi, wo)
	pdf := f.Pdf(wi, wo)
	return fr, wi, pdf, (BSDF_DIFFUSE | BSDF_REFLECTION)
}

func (f *LambertReflection) Type() BsdfType {
	return BSDF_DIFFUSE | BSDF_REFLECTION
}
