package core

type Vector2d struct {
    X, Y Float
}

func NewVector2d(x, y Float) *Vector2d {
    v := &Vector2d{}
    v.X = x
    v.Y = y
    return v
}
