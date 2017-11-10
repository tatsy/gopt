package core

type Accelerator interface {
    Intersect(ray *Ray, isect *Intersection) bool
}

type Shape interface {
    Intersect(ray *Ray, tHit *Float, isect *Intersection) bool
    SampleP(u Point2d, pos *Vector3d, normal *Vector3d)
    Bounds() *Bounds3d
}

type Bxdf interface {
    Pdf(wi, wo *Vector3d) Float
    Sample(wo *Vector3d, u *Point2d) (*Color, *Vector3d, Float)
    Type() int
}

type Sensor interface {
    Film() *Film
    SpawnRay(x, y int) *Ray
}

type Sampler interface {
    Get1D() Float
    Get2D() *Point2d
}

type Light interface {
    Le() *Color
    LeWithRay(ray *Ray) *Color
}
