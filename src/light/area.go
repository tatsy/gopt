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

func (l *AreaLight) Le() *Color {
    return l.radiance
}

func (l *AreaLight) LeWithRay(ray *Ray) *Color {
    var tHit Float
    var isect Intersection
    if l.shape.Intersect(ray, &tHit, &isect) {
        dot := -ray.Dir.Dot(isect.Normal)
        return l.radiance.Scale(dot / (tHit * tHit))
    }
    return NewColor(0.0, 0.0, 0.0)
}
