package core

type Accelerator interface {
    Intersect(ray *Ray, isect *Intersection) bool
}

type Shape interface {
    Intersect(ray *Ray, isect *Intersection) bool
    SampleP(u Point2d, pos *Vector3d, normal *Vector3d)
    Bounds() *Bounds3d
}

type Bsdf interface {
    Pdf(wi, wo Vector3d) Float
    Sample(wo Vector3d, sampler Sampler) Vector3d
}

type Sensor interface {
    Film() *Film
    SpawnRay(x, y int) *Ray
}

type Sampler interface {
    Get1D() Float
    Get2D() Point2d
}

type Light interface {
    Le(ray *Ray) *Color
}
