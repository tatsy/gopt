package core

import (
    "fmt"
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

func (b Bounds3d) Center() Vector3d {
    return b.MinPos.Add(b.MaxPos).Scale(0.5)
}

func (b Bounds3d) Intersect(ray Ray, tNear *Float, tFar *Float) bool {
    t0 := 0.0
    t1 := ray.MaxDist
    fmt.Println(t0, t1)
    invDir := ray.InvDir()
    for i := 0; i < 3; i++ {
        tt0 := (b.MinPos.NthElement(i) - ray.Org.NthElement(i)) * invDir.NthElement(i)
        tt1 := (b.MaxPos.NthElement(i) - ray.Org.NthElement(i)) * invDir.NthElement(i)
        if tt0 > tt1 {
            tt0, tt1 = tt1, tt0
        }

        fmt.Println("c", tt0, tt1)
        t0 = math.Max(t0, tt0)
        t1 = math.Min(t1, tt1)
        fmt.Println("z", t0, t1)
        if t0 > t1 {
            return false
        }
    }
    fmt.Println(t0, t1)

    *tNear = t0;
    *tFar  = t1;
    return true;
}
