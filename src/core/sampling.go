package core

import (
    "math"
)

func CoordinateSystem(w *Vector3d) (*Vector3d, *Vector3d) {
    var u, v *Vector3d
    if math.Abs(w.X) > 0.1 {
        u = NewVector3d(0.0, 1.0, 0.0).Cross(w).Normalized()
    } else {
        u = NewVector3d(1.0, 0.0, 0.0).Cross(w).Normalized()
    }
    v = w.Cross(u).Normalized()
    return u, v
}
