package core

import (
    "testing"
)

func TestNewBounds3d(t *testing.T) {
    b := NewBounds3d()
    if b.MinPos.X != Infinity || b.MinPos.Y != Infinity || b.MinPos.Z != Infinity ||
       b.MaxPos.X != -Infinity || b.MaxPos.Y != -Infinity || b.MaxPos.Z != -Infinity {
        t.Error("Initialization failed")
    }
}

func TestMerge(t *testing.T) {
    b1 := NewBounds3dMinMax(
        NewVector3d(0.0, 0.0, 0.0),
        NewVector3d(1.0, 1.0, 1.0),
    )

    b2 := NewBounds3dMinMax(
        NewVector3d(0.5, 0.5, 0.5),
        NewVector3d(2.0, 2.0, 2.0),
    )

    b3 := b1
    b3.Merge(b2)
    if !b3.MinPos.Equals(b1.MinPos) && !b3.MaxPos.Equals(b2.MaxPos) {
        t.Error("Failed")
    }
}

func TestBounds3dIntersect(t *testing.T) {
    b1 := NewBounds3dMinMax(
        NewVector3d(0.0, 0.0, 0.0),
        NewVector3d(1.0, 1.0, 1.0),
    )
    r := NewRay(
        NewVector3d(1.5, 0.5, 0.5),
        NewVector3d(-1.0, 0.0, 0.0),
    )

    var tMin, tMax Float
    if !b1.Intersect(r, &tMin, &tMax) {
        t.Error("Failed")
    }

    if tMin != 0.5 || tMax != 1.5 {
        t.Errorf("Failed: %f vs %f", tMin, 0.5)
    }
}

func TestMaxExtent(t *testing.T) {
    b1 := NewBounds3dMinMax(
        NewVector3d(1.0, 2.0, 3.0),
        NewVector3d(2.0, 4.0, 6.0),
    )

    if b1.MaxExtent() != 2 {
        t.Error("Max extent test is failed!")
    }
}
