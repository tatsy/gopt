package core

import (
    "math"
)

type Ray struct {
    Org *Vector3d
    Dir *Vector3d
    MaxDist Float
}

func NewRay(org *Vector3d, dir *Vector3d) *Ray {
    return &Ray{
        Org: org,
        Dir: dir,
        MaxDist: Infinity,
    }
}

func NewRayBetweenPoints(origin, target *Vector3d) *Ray {
    dir := target.Subtract(origin)
    dist := dir.Length()
    dir = dir.Divide(dist)

    ray := &Ray{}
    ray.Org = origin.Add(dir.Scale(Eps))
    ray.Dir = dir
    ray.MaxDist = dist - 2.0 * Eps
    return ray
}

func (ray *Ray) Clone() *Ray {
    return &Ray{
        Org: ray.Org,
        Dir: ray.Dir,
        MaxDist: ray.MaxDist,
    }
}

func (r *Ray) InvDir() *Vector3d {
    d := &Vector3d{}
    if (math.Abs(r.Dir.X) > Eps) {
        d.X = 1.0 / r.Dir.X
    } else {
        d.X = Infinity * Sign(r.Dir.X)
    }

    if (math.Abs(r.Dir.Y) > Eps) {
        d.Y = 1.0 / r.Dir.Y
    } else {
        d.Y = Infinity * Sign(r.Dir.Y)
    }

    if (math.Abs(r.Dir.Z) > Eps) {
        d.Z = 1.0 / r.Dir.Z
    } else {
        d.Z = Infinity * Sign(r.Dir.Z)
    }
    return d
}
