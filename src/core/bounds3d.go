package core

import (
    "math"
)

type Bounds3d struct {
    MinPos, MaxPos Vector3d
}

func NewBounds3d() (b Bounds3d) {
    b.MinPos = Vector3d{Infinity, Infinity, Infinity}
    b.MaxPos = Vector3d{-Infinity, -Infinity, -Infinity}
    return
}

func (b1 *Bounds3d) Merge(b2 Bounds3d) {
    b1.MinPos = b1.MinPos.Minimum(b2.MinPos)
    b1.MaxPos = b1.MaxPos.Maximum(b2.MaxPos)
    return
}

func (b *Bounds3d) Center() (c Vector3d) {
    c = b.MinPos.Add(b.MaxPos)
    c = c.Scale(0.5)
    return
}

func (b *Bounds3d) Intersect(ray Ray, tNear *Float, tFar *Float) bool {
    t0 := 0.0
    t1 := ray.MaxDist
    invDir := ray.InvDir()
    for i := 0; i < 3; i++ {
        tt0 := (b.MinPos.NthElement(i) - ray.Org.NthElement(i)) * invDir.NthElement(i)
        tt1 := (b.MaxPos.NthElement(i) - ray.Org.NthElement(i)) * invDir.NthElement(i)
        if tt0 > tt1 {
            tt0, tt1 = tt1, tt0
        }

        t0 = math.Max(t0, tt0)
        t1 = math.Min(t1, tt1)
        if t0 > t1 {
            return false
        }
    }

    *tNear = t0;
    *tFar  = t1;
    return true;
}
