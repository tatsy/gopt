package sampler

import (
	"math/rand"

	. "github.com/tatsy/gopt/src/core"
)

// IndependentSampler generates pseudo random numbers
// with Go's buildin random number generator.
type IndependentSampler struct {
}

func (sampler *IndependentSampler) Get1D() Float {
	return Float(rand.Float64())
}

func (sampler *IndependentSampler) Get2D() *Vector2d {
	return NewVector2d(rand.Float64(), rand.Float64())
}
