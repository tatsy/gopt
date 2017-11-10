package core

const (
    BSDF_DIFFUSE int = 1 << iota
    BSDF_SPECULAR
    BSDF_GLOSSY
    BSDF_REFLECTION
    BSDF_TRANSMISSION
)

type Bsdf struct {
    ns, ts, bs *Vector3d
    bxdf Bxdf
}

func NewBsdf(isect *Intersection, bxdf Bxdf) *Bsdf {
    bsdf := &Bsdf{}
    bsdf.ns = isect.Normal
    bsdf.ts, bsdf.bs = CoordinateSystem(bsdf.ns)
    bsdf.bxdf = bxdf
    return bsdf
}

func (bsdf *Bsdf) SampleWi(wo *Vector3d, u *Point2d) (*Color, *Vector3d, Float, int) {
    xLocal := bsdf.ts.Dot(wo)
    yLocal := bsdf.bs.Dot(wo)
    zLocal := bsdf.ns.Dot(wo)
    woLocal := NewVector3d(xLocal, yLocal, zLocal)
    f, wiLocal, pdf := bsdf.bxdf.Sample(woLocal, u)
    bsdfType := bsdf.bxdf.Type()
    wi := bsdf.ts.Scale(wiLocal.X).
          Add(bsdf.bs.Scale(wiLocal.Y)).
          Add(bsdf.ns.Scale(wiLocal.Z))
    return f, wi, pdf, bsdfType
}
