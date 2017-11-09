package bsdf

import (
    // "math"
    . "../core"
)

type LambertBsdf struct {
    Re Color
}

func (bsdf *LambertBsdf) Pdf(wi, wo Vector3d) Float {
    return 1.0
}

func (bsdf *LambertBsdf) Sample(wo Vector3d, sampler Sampler) (wi Vector3d) {
    // phi := 2.0 * math.Pi * sampler.Get()
    // z2 := sampler.Get()
    return Vector3d{}
}
