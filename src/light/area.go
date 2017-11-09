package light

import (
    . "core"
)

type AreaLight struct {
    shape Shape
    radiance *Color
}

func NewAreaLight(shape Shape, radiance *Color) *AreaLight {
    a := &AreaLight{}
    a.shape = shape
    a.radiance = radiance
    return a
}

func (l *AreaLight) SampleP(u Point2d, pos *Vector3d, normal *Vector3d, L *Color) {
    L = l.radiance
    l.shape.SampleP(u, pos, normal)
}

func (l *AreaLight) Le(ray *Ray) *Color {
    var isect Intersection
    if l.shape.Intersect(ray, &isect) {
        dot := -ray.Dir.Dot(isect.Normal)
        temp := ray.Org.Subtract(isect.Pos)
        dist2 := temp.LengthSquared()
        return l.radiance.Scale(dot / dist2)
    }
    return NewColor(0.0, 0.0, 0.0)
}
