package light

import (
    "math"
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

func (l *AreaLight) SampleLi(isect *Intersection, u *Point2d) (*Color, *Vector3d, Float, *VisibilityTester){
    pos, normal, pdf := l.shape.SampleP(u)
    vt := NewVisibilityTester(isect.Pos, pos)

    wi := pos.Subtract(isect.Pos)
    distSquared := wi.LengthSquared()

    wi = wi.Normalized()
    dot0 := math.Max(0.0, wi.Dot(isect.Normal))
    dot1 := math.Max(0.0, -wi.Dot(normal))

    Le := l.radiance.Scale(dot0 * dot1 / distSquared)
    return Le, wi, pdf, vt
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
