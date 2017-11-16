package core

type BsdfType int

const (
	BSDF_DIFFUSE BsdfType = 1 << iota
	BSDF_SPECULAR
	BSDF_GLOSSY
	BSDF_REFLECTION
	BSDF_TRANSMISSION
)

type Bsdf struct {
	ns, ts, bs *Vector3d
	bxdf       Bxdf
}

func NewBsdf(isect *Intersection, bxdf Bxdf) *Bsdf {
	bsdf := &Bsdf{}
	bsdf.ns = isect.Normal
	bsdf.ts, bsdf.bs = CoordinateSystem(bsdf.ns)
	bsdf.bxdf = bxdf
	return bsdf
}

func (bsdf *Bsdf) Eval(wi, wo *Vector3d) *Color {
	wiLocal := bsdf.toLocal(wi)
	woLocal := bsdf.toLocal(wo)
	return bsdf.bxdf.Eval(wiLocal, woLocal)
}

func (bsdf *Bsdf) Pdf(wi, wo *Vector3d) Float {
	wiLocal := bsdf.toLocal(wi)
	woLocal := bsdf.toLocal(wo)
	return bsdf.bxdf.Pdf(wiLocal, woLocal)
}

func (bsdf *Bsdf) SampleWi(wo *Vector3d, u *Vector2d) (*Color, *Vector3d, Float, BsdfType) {
	woLocal := bsdf.toLocal(wo)
	f, wiLocal, pdf, bsdfType := bsdf.bxdf.Sample(woLocal, u)
	wi := bsdf.ts.Scale(wiLocal.X).
		Add(bsdf.bs.Scale(wiLocal.Y)).
		Add(bsdf.ns.Scale(wiLocal.Z)).
		Normalized()
	return f, wi, pdf, bsdfType
}

func (bsdf *Bsdf) toLocal(wWorld *Vector3d) *Vector3d {
	x := bsdf.ts.Dot(wWorld)
	y := bsdf.bs.Dot(wWorld)
	z := bsdf.ns.Dot(wWorld)
	return NewVector3d(x, y, z).Normalized()
}

func (bsdf *Bsdf) IsSpecular() bool {
	return (bsdf.bxdf.Type() & BSDF_SPECULAR) != 0
}
