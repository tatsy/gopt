package core

import (
    "fmt"
    "math"
    "strings"
)

type Float = float64

const (
    Eps = Float(1.0e-12)
    Infinity = Float(1.0e20)
)

type Semaphore struct {}

func Sign(x Float) Float {
    if x < 0.0 {
        return -1.0
    }
    return 1.0
}

func Clamp(x, lower, upper Float) Float {
    return math.Max(lower, math.Min(x, upper))
}

func ProgressBar(x, maxVal int) {
    width := 50
    ratio := float64(x) / float64(maxVal)
    percent := 100.0 * ratio

    barLen := int(float64(width) * ratio)
    bar := []rune(strings.Repeat("=", barLen) + strings.Repeat(" ", width - barLen))
    if barLen < width {
        bar[barLen] = '>'
    }
    fmt.Printf("\r[ %6.2f %% ] [ %s ]", percent, string(bar))
}
