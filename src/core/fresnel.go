package core

import (
    "math"
)

func FresnelDielectric(cosThetaI, etaI, etaT Float) Float {
    cosThetaI = Clamp(cosThetaI, -1.0, 1.0)

    entering := cosThetaI > 0.0
    if !entering {
        etaI, etaT = etaT, etaI
        cosThetaI = math.Abs(cosThetaI)
    }

    sinThetaI := math.Sqrt(math.Max(0.0, 1.0 - cosThetaI * cosThetaI))
    sinThetaT := etaI / etaT * sinThetaI
    if sinThetaT >= 1.0 {
        return 1.0
    }

    cosThetaT := math.Sqrt(math.Max(0.0, 1.0 - sinThetaT * sinThetaT))
    Rpara := ((etaT * cosThetaI) - (etaI * cosThetaT)) / ((etaT * cosThetaI) + (etaI * cosThetaT))
    Rperp := ((etaI * cosThetaI) - (etaT * cosThetaT)) / ((etaI * cosThetaI) + (etaT * cosThetaT))
    return 0.5 * (Rpara * Rpara + Rperp * Rperp)
}
