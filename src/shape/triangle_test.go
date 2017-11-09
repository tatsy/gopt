package shape

import (
    "testing"
    . "core"
)

func TestTriangleIntersect(t *testing.T) {
    tri := NewTriangle(
        [3]*Vector3d{
            NewVector3d(0.0, 0.0, 0.0),
            NewVector3d(0.0, 1.0, 0.0),
            NewVector3d(1.0, 0.0, 0.0),
        },
        [3]*Vector3d{
            NewVector3d(0.0, 0.0, 1.0),
            NewVector3d(0.0, 0.0, 1.0),
            NewVector3d(0.0, 0.0, 1.0),
        },
    )
    r1 := NewRay(
        NewVector3d(0.5, 0.5, 1.0),
        NewVector3d(0.0, 0.0, -1.0),
    )

    var isect Intersection
    if !tri.Intersect(r1, &isect) {
        t.Error("Ray must intersect, but not intersected!")
    }

    r2 := NewRay(
        NewVector3d(2.0, 2.0, 1.0),
        NewVector3d(0.0, 0.0, -1.0),
    )

    if tri.Intersect(r2, &isect) {
        t.Error("Ray must not intersect, but intersected!")
    }
}
