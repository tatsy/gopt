package core

import (
	"math"
	"testing"
	"time"
)

func TestRandomFloat64(t *testing.T) {
	r := NewRandom(time.Now().UnixNano())
	for trial := 0; trial < TestTrials; trial++ {
		v := r.Float64()
		if v < 0.0 || v > 1.0 {
			t.Errorf("Float64() returns value out of [0, 1]")
		}
	}

	count := 0
	numSamples := 10000000

	for trial := 0; trial < numSamples; trial++ {
		vx := r.Float64()
		vy := r.Float64()
		if vx*vx+vy*vy < 1.0 {
			count++
		}
	}

	estimated := (4.0 * Float(count)) / Float(numSamples)
	if math.Abs(estimated-math.Pi) > 1.0e-3 {
		t.Errorf("Monte carlo Pi is invalid: %f", estimated)
	}
}
