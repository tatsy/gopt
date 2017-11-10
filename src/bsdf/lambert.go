package bsdf

import (
    "math"
    . "core"
)

type LambertReflection struct {
    re *Color
}

func NewLambertReflection(re *Color) *LambertReflection {
    f := &LambertReflection{}
    f.re = re
    return f
}

func (f *LambertReflection) Pdf(wi, wo *Vector3d) Float {
    return 1.0
}

func (f *LambertReflection) Sample(wo *Vector3d, u *Point2d) (*Color, *Vector3d, Float) {
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
