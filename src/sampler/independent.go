package sampler

import (
    "math/rand"
    . "core"
)

type IndependentSampler struct {
}

func (sampler *IndependentSampler) Get() Float {
    return Float(rand.Float64())
}
