package core

import (
	"math"
)

func CosTheta(w *Vector3d) Float {
	return w.Z
}

func AbsCosTheta(w *Vector3d) Float {
	return math.Abs(CosTheta(w))
}

func SameHemisphere(wi, wo *Vector3d) bool {
	return wi.Z*wo.Z > 0.0
}

func Refract(wi, n *Vector3d, eta Float) (*Vector3d, bool) {
	cosThetaI := n.Dot(wi)
	sin2ThetaI := math.Max(0.0, 1.0-cosThetaI*cosThetaI)
	sin2ThetaT := eta * eta * sin2ThetaI
	if sin2ThetaT >= 1.0 {
		return nil, false
	}

	cosThetaT := math.Sqrt(1.0 - sin2ThetaT)
	wt := wi.Scale(-eta).Add(n.Scale(eta*cosThetaI - cosThetaT))
	wt = wt.Normalized()
	return wt, true
}
