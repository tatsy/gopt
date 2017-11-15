package core

import (
	"math"
	"testing"
)

func TestTransformAt(t *testing.T) {
	m := NewTransform([4][4]Float{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			expected := i*4 + j + 1
			if m.At(i, j) != Float(expected) {
				t.Errorf("m.At(%d, %d) != %f, detected %d", i, j, expected, m.At(i, j))
			}
		}
	}

	defer func() {
		if p := recover(); p != nil {
		}
	}()

	m.At(-1, -1)
	t.Error("Out of bounds subscription did not panic.")
}

func TestTransformInverted(t *testing.T) {
	m := NewTransform([4][4]Float{
		{1, 2, 3, 4},
		{2, 2, 3, 4},
		{3, 3, 3, 4},
		{4, 4, 4, 4},
	})

	mi := m.Multiply(m.Inverted())
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				if math.Abs(mi.At(i, j)-1.0) >= Eps {
					t.Errorf("Matrix inversion failed!")
				}
			} else {
				if math.Abs(mi.At(i, j)) >= Eps {
					t.Errorf("Matrix inversion failed!")
				}
			}
		}
	}
}
