package core

import (
	"math"
)

// Ray is exhibited from the position Org to the direction Dir.
type Ray struct {
	Org     *Vector3d
	Dir     *Vector3d
	InvDir  *Vector3d
	MaxDist Float
}

func NewRay(org *Vector3d, dir *Vector3d) *Ray {
	r := new(Ray)
	r.Org = org
	r.Dir = dir
	r.MaxDist = Infinity
	r.InvDir = invertDir(r.Dir)
	return r
}

func NewRayBetweenPoints(origin, target *Vector3d) *Ray {
	dir := target.Subtract(origin)
	dist := dir.Length()
	dir = dir.Divide(dist)

	ray := &Ray{}
	ray.Org = origin.Add(dir.Scale(Eps))
	ray.Dir = dir
	ray.InvDir = invertDir(dir)
	ray.MaxDist = dist - 2.0*Eps
	return ray
}

func (ray *Ray) Clone() *Ray {
	return &Ray{
		Org:     ray.Org,
		Dir:     ray.Dir,
		InvDir:  ray.InvDir,
		MaxDist: ray.MaxDist,
	}
}

func invertDir(v *Vector3d) *Vector3d {
	d := &Vector3d{}
	if math.Abs(v.X) > Eps {
		d.X = 1.0 / v.X
	} else {
		d.X = Infinity * Sign(v.X)
	}

	if math.Abs(v.Y) > Eps {
		d.Y = 1.0 / v.Y
	} else {
		d.Y = Infinity * Sign(v.Y)
	}

	if math.Abs(v.Z) > Eps {
		d.Z = 1.0 / v.Z
	} else {
		d.Z = Infinity * Sign(v.Z)
	}
	return d
}
