package core

import (
	"math/rand"
	"testing"
)

func NewVector2dTest(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		x := rand.Float64()
		y := rand.Float64()
		v := NewVector2d(x, y)
		if v.X != x || v.Y != y {
			t.Errorf("Expected: (%f, %f), Actual: (%f, %f)", x, y, v.X, v.Y)
		}
	}
}

func Vector2dNthElementTest(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		x := rand.Float64()
		y := rand.Float64()
		v := NewVector2d(x, y)
		if x != v.NthElement(0) {
			t.Errorf("%d-th element mismatched: Expected: %f, Actual: %f", 0, x, v.X)
		}

		if y != v.NthElement(1) {
			t.Errorf("%d-th element mismatched: Expected: %f, Actual: %f", 0, y, v.Y)
		}

		defer func() {
			if r := recover(); r != nil {
			}
		}()

		var k int
		for 0 <= k && k <= 1 {
			k = rand.Int()
		}
		v.NthElement(k)
		t.Errorf("Out of bounds element access did not panic")
	}
}
