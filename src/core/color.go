package core

import (
    "fmt"
)

type Color struct {
    R, G, B Float
}

func NewColor(r, g, b Float) *Color {
    c := &Color{}
    c.R = r
    c.G = g
    c.B = b
    return c
}

func NewColorWithString(s string) *Color {
    var r, g, b Float
    n, _ := fmt.Sscanf(s, "(%f, %f, %f)", &r, &g, &b)
    if n != 3 {
        panic(fmt.Sprintf("Failed to parse Vector3d: %s", s))
    }
    return &Color{r, g, b}
}


func (c *Color) Y() Float {
    return 0.299 * c.R + 0.587 * c.G + 0.114 * c.B
}

func (c *Color) IsBlack() bool {
    return c.R == 0.0 && c.G == 0.0 && c.B == 0.0
}

func (c1 *Color) Add(c2 *Color) *Color {
    ret := &Color{}
    ret.R = c1.R + c2.R
    ret.G = c1.G + c2.G
    ret.B = c1.B + c2.B
    return ret
}

func (c *Color) Scale(x Float) *Color {
    ret := &Color{}
    ret.R = c.R * x
    ret.G = c.G * x
    ret.B = c.B * x
    return ret
}

func (c1 *Color) Multiply(c2 *Color) *Color {
    ret := &Color{}
    ret.R = c1.R * c2.R
    ret.G = c1.G * c2.G
    ret.B = c1.B * c2.B
    return ret
}
