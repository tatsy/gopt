package core

type Primitive struct {
    Shape Shape
    Bsdf Bsdf
}

func NewPrimitive(shape Shape, bsdf Bsdf) (prim Primitive) {
    prim.Shape = shape
    prim.Bsdf = bsdf
    return
}
