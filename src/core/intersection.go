package core

type Intersection struct {
    Pos, Normal *Vector3d
    HitDist Float
}

func NewIntersection(pos, normal *Vector3d, tHit Float) *Intersection {
    isect := &Intersection{}
    isect.Pos = pos
    isect.Normal = normal
    isect.HitDist = tHit
    return isect
}
