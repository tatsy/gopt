package core

type Intersection struct {
    Pos, Normal *Vector3d
    primitive *Primitive
}

func NewIntersection(pos, normal *Vector3d) *Intersection {
    isect := &Intersection{}
    isect.Pos = pos
    isect.Normal = normal
    isect.primitive = nil
    return isect
}

func (isect *Intersection) SpawnRay(wi *Vector3d) *Ray {
    return NewRay(isect.Pos, wi)
}

func (isect *Intersection) Bsdf() *Bsdf {
    return NewBsdf(isect, isect.primitive.Bxdf())
}

func (isect *Intersection) Le(wo *Vector3d) *Color {
    if isect.primitive.Light() == nil {
        return NewColor(0.0, 0.0, 0.0)
    }
    dot := wo.Dot(isect.Normal)
    return isect.primitive.Light().Le().Scale(dot)
}
