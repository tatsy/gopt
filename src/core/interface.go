package core

type Shape interface {
    Intersect(ray Ray, isect *Intersection) bool
    Bounds() Bounds3d
}

type Bsdf interface {
    Pdf(wi, wo Vector3d) Float
    Sample(wo Vector3d, sampler Sampler) Vector3d
}


type Sensor interface {
    Film() Film
    SpawnRay(x, y int) Ray
}

type Sampler interface {
    Get() Float
}
