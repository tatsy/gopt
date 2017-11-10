package shape

import (
    "math"
    . "core"
)

type Triangle struct {
    Points [3]*Vector3d
    Normals [3]*Vector3d
}

func NewTriangle(points, normals [3]*Vector3d) *Triangle {
    t := &Triangle{}
    t.Points = points
    t.Normals = normals
    return t
}

func (t *Triangle) Intersect(ray *Ray, tHit *Float, isect *Intersection) bool {
    e1 := t.Points[1].Subtract(t.Points[0]);
    e2 := t.Points[2].Subtract(t.Points[0]);
    pVec := ray.Dir.Cross(e2)
    det := e1.Dot(pVec)
    if det > -Eps && det < Eps {
        return false
    }

    invDet := 1.0 / det
    tVec := ray.Org.Subtract(t.Points[0])
    u := tVec.Dot(pVec) * invDet
    if u < 0.0 || u > 1.0 {
        return false
    }

    qVec := tVec.Cross(e1)
    v := ray.Dir.Dot(qVec) * invDet
    if v < 0.0 || u + v > 1.0 {
        return false
    }

    *tHit = e2.Dot(qVec) * invDet
    if *tHit <= Eps || *tHit > ray.MaxDist {
        return false
    }

    pos := ray.Org.Add(ray.Dir.Scale(*tHit))
    nrm := t.Normals[0].Scale(1.0 - u - v).
           Add(t.Normals[1].Scale(u)).
           Add(t.Normals[2].Scale(v))
    //Point2d uv = (1.0 - u - v) * uvs_[0] + u * uvs_[1] + v * uvs_[2];

    *isect = *NewIntersection(pos, nrm, ray.Dir.Negate())
    return true
}

func (t *Triangle) SampleP(rnd *Point2d) (*Vector3d, *Vector3d, Float) {
    u, v := rnd.X, rnd.Y
    if u + v >= 1.0 {
        u = 1.0 - u
        v = 1.0 - v
    }

    pos := t.Points[0].Scale(1.0 - u - v).
           Add(t.Points[1].Scale(u)).
           Add(t.Points[2].Scale(v))
    normal := t.Normals[0].Scale(1.0 - u - v).
              Add(t.Normals[1].Scale(u)).
              Add(t.Normals[2].Scale(v))
    area := 0.5 * (t.Points[1].Subtract(t.Points[0])).Cross(t.Points[2].Subtract(t.Points[0])).Length()
    pdf := 0.0
    if area != 0.0 {
        pdf = 1.0 / area
    }
    return pos, normal, pdf
}

func (t *Triangle) Bounds() *Bounds3d {
    b := NewBounds3d()
    b.MinPos.X = math.Min(t.Points[0].X, math.Min(t.Points[1].X, t.Points[2].X))
    b.MinPos.Y = math.Min(t.Points[0].Y, math.Min(t.Points[1].Y, t.Points[2].Y))
    b.MinPos.Z = math.Min(t.Points[0].Z, math.Min(t.Points[1].Z, t.Points[2].Z))
    b.MaxPos.X = math.Max(t.Points[0].X, math.Max(t.Points[1].X, t.Points[2].X))
    b.MaxPos.Y = math.Max(t.Points[0].Y, math.Max(t.Points[1].Y, t.Points[2].Y))
    b.MaxPos.Z = math.Max(t.Points[0].Z, math.Max(t.Points[1].Z, t.Points[2].Z))
    return b
}
