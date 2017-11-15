package light

import (
	"math"

	. "github.com/tatsy/gopt/src/core"
)

// AreaLight represents an area light
type AreaLight struct {
	shape    Shape  // Shape of the area light
	radiance *Color // Radiance of this area light
}

func NewAreaLight(shape Shape, radiance *Color) *AreaLight {
	a := &AreaLight{}
	a.shape = shape
	a.radiance = radiance
	return a
}

func (l *AreaLight) SampleLi(isect *Intersection, u *Vector2d) (*Color, *Vector3d, Float, *VisibilityTester) {
	pos, normal := l.shape.SamplePoint(u)
	vt := NewVisibilityTester(isect.Pos, pos)

	wi := pos.Subtract(isect.Pos).Normalized()
	pdf := l.shape.Pdf(isect, wi)

	dot := math.Max(0.0, -wi.Dot(normal))
	Le := l.radiance
	if dot <= 0.0 {
		Le = NewColor(0.0, 0.0, 0.0)
	}
	return Le, wi, pdf, vt
}

func (l *AreaLight) PdfLi(isect *Intersection, wi *Vector3d) Float {
	return l.shape.Pdf(isect, wi)
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
