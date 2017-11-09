package core

import (
    "math/rand"
    "testing"
)

func TestNewColor(t *testing.T) {
    r := rand.Float64()
    g := rand.Float64()
    b := rand.Float64()
    c := NewColor(r, g, b)
    if c.R != r || c.G != g || c.B != b {
        t.Errorf("Color initialization failed: %v != (%f, %f, %f)", c, r, g, b)
    }
}

func TestColorY(t *testing.T) {
    r := rand.Float64()
    g := rand.Float64()
    b := rand.Float64()
    c := NewColor(r, g, b)
    actual := c.Y()
    expected := 0.299 * c.R + 0.587 * c.G + 0.114 * c.B
    if actual != expected {
        t.Errorf("Color.Y: expected %v, actual %v", expected, actual)
    }
}
