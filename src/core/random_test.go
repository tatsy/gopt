package core

import (
	"testing"
	"time"
)

const (
	numTrials = 100
)

func TestRandomFloat64(t *testing.T) {
	r := NewRandom(time.Now().UnixNano())
	for trial := 0; trial < numTrials; trial++ {
		v := r.Float64()
		if v < 0.0 || v > 1.0 {
			t.Errorf("Float64() returns value out of [0, 1]")
		}
	}
}
