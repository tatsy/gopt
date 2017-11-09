package core

type Primitive struct {
    Shape Shape
    Bsdf Bsdf
    Light Light
}

func NewPrimitive(shape Shape, bsdf Bsdf) *Primitive {
    prim := &Primitive{}
    prim.Shape = shape
    prim.Bsdf = bsdf
    prim.Light = nil
    return prim
}

func NewLightPrimitive(shape Shape, bsdf Bsdf, light Light) *Primitive {
    prim := &Primitive{}
    prim.Shape = shape
    prim.Bsdf = bsdf
    prim.Light = light
    return prim
}
