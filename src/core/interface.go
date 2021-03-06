package core

type Accelerator interface {
	Intersect(ray *Ray, isect *Intersection) bool
}

type Shape interface {
	Intersect(ray *Ray, tHit *Float, isect *Intersection) bool
	SamplePoint(u *Vector2d) (*Vector3d, *Vector3d)
	Pdf(isect *Intersection, wi *Vector3d) Float
	Bounds() *Bounds3d
}

type Bxdf interface {
	Pdf(wi, wo *Vector3d) Float
	Eval(wi, wo *Vector3d) *Color
	Sample(wo *Vector3d, u *Vector2d) (*Color, *Vector3d, Float, BsdfType)
	Type() BsdfType
}

type Sensor interface {
	Film() *Film
	SpawnRay(x, y Float, u *Vector2d) *Ray
}

type Sampler interface {
	Clone(seed int64) Sampler
	Get1D() Float
	Get2D() *Vector2d
}

type Light interface {
	Le() *Color
	LeWithRay(ray *Ray) *Color
	SampleLi(isect *Intersection, u *Vector2d) (*Color, *Vector3d, Float, *VisibilityTester)
	PdfLi(isect *Intersection, wi *Vector3d) Float
}
