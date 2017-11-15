package sampler

import (
	. "github.com/tatsy/gopt/src/core"
)

// IndependentSampler generates pseudo random numbers
// with Go's buildin random number generator.
type IndependentSampler struct {
	*Random
}

func NewIndependentSampler() *IndependentSampler {
	s := new(IndependentSampler)
	s.Random = NewRandom(0)
	return s
}

func (*IndependentSampler) Clone(seed int64) Sampler {
	s := new(IndependentSampler)
	s.Random = NewRandom(seed)
	return s
}

func (s *IndependentSampler) Get1D() Float {
	return Float(s.Float64())
}

func (s *IndependentSampler) Get2D() *Vector2d {
	return NewVector2d(s.Float64(), s.Float64())
}
