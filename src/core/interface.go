package core

type Accelerator interface {
	Intersect(ray *Ray, isect *Intersection) bool
}

type Shape interface {
	Intersect(ray *Ray, tHit *Float, isect *Intersection) bool
	SampleP(u *Vector2d) (*Vector3d, *Vector3d, Float)
	Bounds() *Bounds3d
}

type Bxdf interface {
	Pdf(wi, wo *Vector3d) Float
	Eval(wi, wo *Vector3d) *Color
	Sample(wo *Vector3d, u *Vector2d) (*Color, *Vector3d, Float)
	Type() int
}

type Sensor interface {
	Film() *Film
	SpawnRay(x, y Float) *Ray
}

type Sampler interface {
	Get1D() Float
	Get2D() *Vector2d
}

type Light interface {
	Le() *Color
	LeWithRay(ray *Ray) *Color
	SampleLi(isect *Intersection, u *Vector2d) (*Color, *Vector3d, Float, *VisibilityTester)
}
