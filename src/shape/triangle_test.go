package shape

import (
    "testing"
    . "core"
)

func TestTriangleIntersect(t *testing.T) {
    tri := NewTriangleWithP(
        [3]*Vector3d{
            NewVector3d(0.0, 0.0, 0.0),
            NewVector3d(0.0, 1.0, 0.0),
            NewVector3d(1.0, 0.0, 0.0),
        },
    )
    r1 := NewRay(
        NewVector3d(0.5, 0.5, 1.0),
        NewVector3d(0.0, 0.0, -1.0),
    )

    var tHit Float
    var isect Intersection
    if !tri.Intersect(r1, &tHit, &isect) {
        t.Error("Ray must intersect, but not intersected!")
    } else if (tHit != 1.0) {
        t.Error("Intersection distance differs: %v != %v", tHit, 1.0)
    }


    r2 := NewRay(
        NewVector3d(2.0, 2.0, 1.0),
        NewVector3d(0.0, 0.0, -1.0),
    )

    if tri.Intersect(r2, &tHit, &isect) {
        t.Error("Ray must not intersect, but intersected!")
    }
}
