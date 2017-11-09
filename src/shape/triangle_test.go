package shape

import (
    "testing"
    . "core"
)

func TestTriangleIntersect(t *testing.T) {
    tri := Triangle{
        Points: [3]Vector3d{
            Vector3d{0.0, 0.0, 0.0},
            Vector3d{0.0, 1.0, 0.0},
            Vector3d{1.0, 0.0, 0.0},
        },
        Normals: [3]Vector3d{
            Vector3d{0.0, 0.0, 0.0},
            Vector3d{0.0, 0.0, 0.0},
            Vector3d{0.0, 0.0, 0.0},
        },
    }
    r := NewRay(
        Vector3d{0.5, 0.5, 1.0},
        Vector3d{0.0, 0.0, -1.0},
    )

    var isect Intersection
    if !tri.Intersect(r, &isect) {
        t.Error("Failed")
    }
}
