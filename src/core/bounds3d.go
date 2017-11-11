package core

import (
    "math"
)

type Bounds3d struct {
    MinPos, MaxPos *Vector3d
}

func NewBounds3d() *Bounds3d {
    b := &Bounds3d{}
    b.MinPos = NewVector3d(Infinity, Infinity, Infinity)
    b.MaxPos = NewVector3d(-Infinity, -Infinity, -Infinity)
    return b
}

func NewBounds3dMinMax(minPos, maxPos *Vector3d) *Bounds3d {
    b := &Bounds3d{}
    b.MinPos = minPos
    b.MaxPos = maxPos
    return b
}

func (b1 *Bounds3d) Merge(b2 *Bounds3d) *Bounds3d {
    ret := &Bounds3d{}
    ret.MinPos = b1.MinPos.Minimum(b2.MinPos)
    ret.MaxPos = b1.MaxPos.Maximum(b2.MaxPos)
    return ret
}

func (b *Bounds3d) MergePoint(v *Vector3d) *Bounds3d {
    ret := &Bounds3d{}
    ret.MinPos = b.MinPos.Minimum(v)
    ret.MaxPos = b.MaxPos.Maximum(v)
    return ret
}

func (b *Bounds3d) Area() Float {
    diff := b.MaxPos.Subtract(b.MinPos)
    return math.Abs(diff.X * diff.Y * diff.Z)
}

func (b *Bounds3d) Centroid() *Vector3d {
    return b.MinPos.Add(b.MaxPos).Scale(0.5)
}

func (b *Bounds3d) MaxExtent() int {
    v := b.MaxPos.Subtract(b.MinPos)
    v = v.Abs()
    switch {
    case v.X > v.Y && v.X > v.Z:
        return 0
    case v.Y > v.Z:
        return 1
    default:
        return 2
    }
}

func (b *Bounds3d) Intersect(ray *Ray, tNear *Float, tFar *Float) bool {
    t0 := 0.0
    t1 := ray.MaxDist
    invDir := ray.InvDir()

    var tt0, tt1 Float

    // X
    tt0 = (b.MinPos.X - ray.Org.X) * invDir.X
    tt1 = (b.MaxPos.X - ray.Org.X) * invDir.X
    if tt0 > tt1 { tt0, tt1 = tt1, tt0 }

    t0 = math.Max(t0, tt0)
    t1 = math.Min(t1, tt1)
    if t0 > t1 { return false }

    // Y
    tt0 = (b.MinPos.Y - ray.Org.Y) * invDir.Y
    tt1 = (b.MaxPos.Y - ray.Org.Y) * invDir.Y
    if tt0 > tt1 { tt0, tt1 = tt1, tt0 }

    t0 = math.Max(t0, tt0)
    t1 = math.Min(t1, tt1)
    if t0 > t1 { return false }

    // Z
    tt0 = (b.MinPos.Z - ray.Org.Z) * invDir.Z
    tt1 = (b.MaxPos.Z - ray.Org.Z) * invDir.Z
    if tt0 > tt1 { tt0, tt1 = tt1, tt0 }

    t0 = math.Max(t0, tt0)
    t1 = math.Min(t1, tt1)
    if t0 > t1 { return false }

    *tNear = t0;
    *tFar  = t1;
    return true;
}
