package core

type Intersection struct {
	Pos, Normal, Wo *Vector3d
	primitive       *Primitive
}

func NewIntersection(pos, normal, wo *Vector3d) *Intersection {
	isect := &Intersection{}
	isect.Pos = pos
	isect.Normal = normal
	isect.Wo = wo
	isect.primitive = nil
	return isect
}

func (isect *Intersection) SpawnRay(wi *Vector3d) *Ray {
	org := offsetRayOrigin(isect.Pos, isect.Normal, wi)
	return NewRay(org, wi)
}

func (isect *Intersection) Bsdf() *Bsdf {
	return NewBsdf(isect, isect.primitive.Bxdf())
}

func (isect *Intersection) IsHitLight(l Light) bool {
	if isect.primitive.Light() == nil {
		return false
	}
	return l == isect.primitive.Light()
}

func (isect *Intersection) Le(wi *Vector3d) *Color {
	if isect.primitive.Light() == nil {
		return NewColor(0.0, 0.0, 0.0)
	}
	dot := wi.Dot(isect.Normal)
	if dot <= 0.0 {
		return NewColor(0.0, 0.0, 0.0)
	}
	return isect.primitive.Light().Le()
}

func offsetRayOrigin(org, n, w *Vector3d) *Vector3d {
	offset := n.Scale(1.0e-3)
	if w.Dot(n) < 0.0 {
		offset = offset.Negate()
	}
	return org.Add(offset)
}
