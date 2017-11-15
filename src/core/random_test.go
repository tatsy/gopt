package core

import (
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
}
