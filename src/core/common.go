package core

import (
	"fmt"
	"math"
	"strings"
)

// Float represetns the common floating point number used in this package.
// Users can change this statement to float32 to use it.
type Float = float64

const (
	// Eps is a very small value.
	Eps = Float(1.0e-12)
	// Infinity represents an inifinite value.
	Infinity = Float(1.0e20)
)

// Semaphore is used to manage for-loop parallelization.
type Semaphore struct{}

// Sign returns the sign of specified value.
func Sign(x Float) Float {
	if x < 0.0 {
		return -1.0
	}
	return 1.0
}

// Clamp truncate the specified value so that it is in [lower, upper].
func Clamp(x, lower, upper Float) Float {
	return math.Max(lower, math.Min(x, upper))
}

// ProgressBar prints out the progress bar on CUI.
func ProgressBar(x, maxVal int) {
	width := 50
	ratio := float64(x) / float64(maxVal)
	percent := 100.0 * ratio

	barLen := int(float64(width) * ratio)
	bar := []rune(strings.Repeat("=", barLen) + strings.Repeat(" ", width-barLen))
	if barLen < width {
		bar[barLen] = '>'
	}
	fmt.Printf("\r[ %6.2f %% ] [ %s ]", percent, string(bar))
}
