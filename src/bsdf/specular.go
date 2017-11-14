package bsdf

import (
	. "github.com/tatsy/gopt/src/core"
)

// Specular reflection
type SpecularReflection struct {
	re *Color
}

func NewSpecularReflection(re *Color) *SpecularReflection {
	f := &SpecularReflection{}
	f.re = re
	return f
}

func (f *SpecularReflection) Eval(wi, wo *Vector3d) *Color {
	if NewVector3d(-wi.X, -wi.Y, wi.Z).Dot(wo) >= 1.0-Eps {
		return f.re
	}
	return NewColor(0.0, 0.0, 0.0)
}

func (f *SpecularReflection) Pdf(wi, wo *Vector3d) Float {
	if NewVector3d(-wi.X, -wi.Y, wi.Z).Dot(wo) >= 1.0-Eps {
		return 1.0
	}
	return 0.0
}

func (f *SpecularReflection) Sample(wo *Vector3d, u *Vector2d) (*Color, *Vector3d, Float) {
	wi := NewVector3d(-wo.X, -wo.Y, wo.Z)
	fr := f.re.Scale(1.0 / AbsCosTheta(wi))
	return fr, wi, 1.0
}

func (f *SpecularReflection) Type() int {
	return BSDF_SPECULAR | BSDF_REFLECTION
}

// Specular fresnel
type SpecularFresnel struct {
	re, tr     *Color
	etaA, etaB Float
}

func NewSpecularFresnel(re, tr *Color, etaA, etaB Float) *SpecularFresnel {
	f := &SpecularFresnel{}
	f.re = re
	f.tr = tr
	f.etaA = etaA
	f.etaB = etaB
	return f
}

func (f *SpecularFresnel) Eval(wi, wo *Vector3d) *Color {
	cosThetaI := CosTheta(wo)
	F := FresnelDielectric(cosThetaI, f.etaA, f.etaB)
	if SameHemisphere(wi, wo) {
		if NewVector3d(-wi.X, -wi.Y, wi.Z).Dot(wo) >= 1.0-Eps {
			return f.re.Scale(F / AbsCosTheta(wi))
		}
	} else {
		entering := CosTheta(wo) > 0.0
		etaI, etaT := f.etaA, f.etaB
		if !entering {
			etaI, etaT = f.etaB, f.etaA
		}

		faceForwardNormal := NewVector3d(0.0, 0.0, 1.0)
		if wo.Z < 0.0 {
			faceForwardNormal.Z = -1.0
		}

		wt, isRefr := Refract(wo, faceForwardNormal, etaI/etaT)
		if isRefr && wt.Dot(wi) > 1.0-Eps {
			return f.tr.Scale((1.0 - F) / AbsCosTheta(wi))
		}
	}
	return NewColor(0.0, 0.0, 0.0)
}

func (f *SpecularFresnel) Pdf(wi, wo *Vector3d) Float {
	cosThetaI := CosTheta(wo)
	F := FresnelDielectric(cosThetaI, f.etaA, f.etaB)
	if SameHemisphere(wi, wo) {
		if NewVector3d(-wi.X, -wi.Y, wi.Z).Dot(wo) >= 1.0-Eps {
			return F
		}
	} else {
		entering := CosTheta(wo) > 0.0
		etaI, etaT := f.etaA, f.etaB
		if !entering {
			etaI, etaT = f.etaB, f.etaA
		}

		faceForwardNormal := NewVector3d(0.0, 0.0, 1.0)
		if wo.Z < 0.0 {
			faceForwardNormal.Z = -1.0
		}

		wt, isRefract := Refract(wo, faceForwardNormal, etaI/etaT)
		if isRefract && wt.Dot(wi) > 1.0-Eps {
			return 1.0 - F
		}
	}
	return 0.0
}

func (f *SpecularFresnel) Sample(wo *Vector3d, u *Vector2d) (*Color, *Vector3d, Float) {
	F := FresnelDielectric(CosTheta(wo), f.etaA, f.etaB)
	if u.X < F {
		// Reflection
		wi := NewVector3d(-wo.X, -wo.Y, wo.Z)
		pdf := F
		fr := f.re.Scale(F / AbsCosTheta(wi))
		return fr, wi, pdf
	} else {
		// Transmission
		entering := CosTheta(wo) > 0.0
		etaI, etaT := f.etaA, f.etaB
		if !entering {
			etaI, etaT = f.etaB, f.etaA
		}

		faceForwardNormal := NewVector3d(0.0, 0.0, 1.0)
		if wo.Z < 0.0 {
			faceForwardNormal.Z = -1.0
		}

		wi, isRefract := Refract(wo, faceForwardNormal, etaI/etaT)
		if isRefract {
			pdf := 1.0 - F
			ft := f.tr.Scale((1.0 - F) * (etaI * etaI) / (etaT * etaT) / AbsCosTheta(wi))
			return ft, wi, pdf
		}
	}
	return NewColor(0.0, 0.0, 0.0), NewVector3d(0.0, 0.0, 1.0), 0.0
}

func (f *SpecularFresnel) Type() int {
	return BSDF_SPECULAR | BSDF_REFLECTION | BSDF_TRANSMISSION
}
