package core

type Float = float64

const (
    Eps = Float(1.0e-12)
    Infinity = Float(1.0e20)
)

func Sign(x Float) Float {
    if x < 0.0 {
        return -1.0
    }
    return 1.0
}
