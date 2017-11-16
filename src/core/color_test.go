package core

import (
	"math/rand"
	"testing"
)

func newRandomColor() *Color {
	c := new(Color)
	c.R = rand.Float64()
	c.G = rand.Float64()
	c.B = rand.Float64()
	return c
}

func TestNewColor(t *testing.T) {
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	c := NewColor(r, g, b)
	if c.R != r || c.G != g || c.B != b {
		t.Errorf("Color initialization failed: %v != (%f, %f, %f)", c, r, g, b)
	}
}

func TestColorMultiply(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		c1 := newRandomColor()
		c2 := newRandomColor()
		c3 := c1.Multiply(c2)
		if c1.R*c2.R != c3.R || c1.G*c2.G != c3.G || c1.B*c2.B != c3.B {
			t.Errorf("%v * %v != %v", c1, c2, c3)
		}
	}
}

func TestColorY(t *testing.T) {
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	c := NewColor(r, g, b)
	actual := c.Y()
	expected := 0.299*c.R + 0.587*c.G + 0.114*c.B
	if actual != expected {
		t.Errorf("Color.Y: expected %v, actual %v", expected, actual)
	}
}
