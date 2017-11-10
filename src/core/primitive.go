package core

type Primitive struct {
    shape Shape
    bxdf Bxdf
    light Light
}

func NewPrimitive(shape Shape, bxdf Bxdf) *Primitive {
    prim := &Primitive{}
    prim.shape = shape
    prim.bxdf = bxdf
    prim.light = nil
    return prim
}

func NewLightPrimitive(shape Shape, bxdf Bxdf, light Light) *Primitive {
    prim := &Primitive{}
    prim.shape = shape
    prim.bxdf = bxdf
    prim.light = light
    return prim
}

func (p *Primitive) Intersect(ray *Ray, isect *Intersection) bool {
    var tHit Float
    if !p.shape.Intersect(ray, &tHit, isect) {
        return false
    }

    ray.MaxDist = tHit
    isect.primitive = p
    return true
}

func (p *Primitive) Bxdf() Bxdf {
    return p.bxdf
}

func (p *Primitive) Light() Light {
    return p.light
}

func (p *Primitive) Bounds() *Bounds3d {
    return p.shape.Bounds()
}
