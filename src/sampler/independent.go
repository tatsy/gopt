package sampler

import (
    "math/rand"
    . "core"
)

type IndependentSampler struct {
}

func (sampler *IndependentSampler) Get1D() Float {
    return Float(rand.Float64())
}

func (sampler *IndependentSampler) Get2D() *Point2d {
    return NewPoint2d(rand.Float64(), rand.Float64())
}
