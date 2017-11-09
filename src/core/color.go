package core

type Color struct {
    R, G, B Float
}

func NewColor(r, g, b Float) (c Color) {
    c.R = r
    c.G = g
    c.B = b
    return
}

func (c *Color) Y() Float {
    return 0.299 * c.R + 0.587 * c.G + 0.114 * c.B
}
