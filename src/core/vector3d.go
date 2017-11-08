package core

import (
    "math"
)

type Vector3d struct {
    X, Y, Z Float
}

func (v1 Vector3d) Add(v2 Vector3d) (ret Vector3d) {
    ret.X = v1.X + v2.X
    ret.Y = v1.Y + v2.Y
    ret.Z = v1.Z + v2.Z
    return
}

func (v1 Vector3d) Negate() (ret Vector3d) {
    ret.X = -v1.X
    ret.Y = -v1.Y
    ret.Z = -v1.Z
    return
}

func (v1 Vector3d) Subtract(v2 Vector3d) (ret Vector3d) {
    ret = v1.Add(v2.Negate())
    return
}

func (v1 Vector3d) Scale(s Float) (ret Vector3d) {
    ret.X = v1.X * s
    ret.Y = v1.Y * s
    ret.Z = v1.Z * s
    return
}

func (v1 Vector3d) Divide(s Float) (ret Vector3d) {
    if s == 0.0 {
        panic("Zero division!")
    }
    ret = v1.Scale(1.0 / s)
    return
}

func (v1 Vector3d) Dot(v2 Vector3d) (ret Float) {
    ret = v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
    return
}

func (v1 Vector3d) Cross(v2 Vector3d) (ret Vector3d) {
    ret.X = v1.Y * v2.Z - v2.Y * v1.Z
    ret.Y = v1.Z * v2.X - v2.Z * v1.X
    ret.Z = v1.X * v2.Y - v2.X * v1.Y
    return
}

func (v1 Vector3d) Length() (ret Float) {
    ret = Float(math.Sqrt(float64(v1.LengthSquared())))
    return
}

func (v1 Vector3d) LengthSquared() (ret Float) {
    ret = v1.Dot(v1)
    return
}

func (v1 Vector3d) Normalized() (ret Vector3d) {
    ret = v1.Divide(v1.Length())
    return
}

func (v1 Vector3d) Minimum(v2 Vector3d) (ret Vector3d) {
    ret.X = math.Min(v1.X, v2.X)
    ret.Y = math.Min(v1.Y, v2.Y)
    ret.Z = math.Min(v1.Z, v2.Z)
    return
}

func (v1 Vector3d) Maximum(v2 Vector3d) (ret Vector3d) {
    ret.X = math.Max(v1.X, v2.X)
    ret.Y = math.Max(v1.Y, v2.Y)
    ret.Z = math.Max(v1.Z, v2.Z)
    return
}

func (v Vector3d) NthElement(i int) Float {
    switch i {
    case 0:
        return v.X
    case 1:
        return v.Y
    case 2:
        return v.Z
    }
    panic("Element index out of range!")
}

func (v1 Vector3d) Equals(v2 Vector3d) bool {
    return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}
