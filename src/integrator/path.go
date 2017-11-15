package integrator

import (
	"fmt"
	"math"

	. "github.com/tatsy/gopt/src/core"
)

type PathIntegrator struct {
}

func (integrator *PathIntegrator) Render(scene *Scene, sensor Sensor, sampler Sampler, params *RenderParams) {
	width := sensor.Film().Width
	height := sensor.Film().Height
	numSamples := params.GetInt("integrator.num-samples")
	maxBounces := params.GetInt("integrator.max-bounces")
	film := sensor.Film()
	sem := make(chan struct{}, width)

	for s := 0; s < numSamples; s++ {
		fmt.Printf("Sample (%d / %d):\n", s+1, numSamples)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				go func(x, y int) {
					seed := int64(s)*int64(width*height) + int64(y*width+x)
					subSampler := sampler.Clone(seed)
					subPos := subSampler.Get2D()
					px := Float(x) + subPos.X
					py := Float(y) + subPos.Y
					ray := sensor.SpawnRay(px, py, subSampler.Get2D())
					L := integrator.Li(scene, ray, subSampler, maxBounces)
					film.AddSample(px, py, L)
					sem <- struct{}{}
				}(x, y)
			}

			for x := 0; x < width; x++ {
				<-sem
			}

			ProgressBar(y+1, height)
		}
		fmt.Println()
		film.Save(params.GetString("outfile"))
	}
}

func (integrator *PathIntegrator) Li(scene *Scene, r *Ray, sampler Sampler, maxBounces int) *Color {
	ray := r.Clone()
	L := NewColor(0.0, 0.0, 0.0)
	beta := NewColor(1.0, 1.0, 1.0)
	specularBounce := false

	for bounces := 0; ; bounces++ {
		var isect Intersection
		isIntersect := scene.Intersect(ray, &isect)
		if bounces == 0 || specularBounce {
			if isIntersect {
				rr := beta.Multiply(isect.Le(ray.Dir.Negate()))
				L = L.Add(rr)
			} else {
				for _, l := range scene.Lights {
					rr := beta.Multiply(l.LeWithRay(ray))
					L = L.Add(rr)
				}
			}
		}

		if !isIntersect || bounces >= maxBounces {
			break
		}

		if isect.Bsdf().IsNotSpecular() {
			Ld := NextEventEstimation(&isect, scene, sampler)
			L = L.Add(beta.Multiply(Ld))
		}

		wo := ray.Dir.Negate()
		f, wi, pdf, bsdfType := isect.Bsdf().SampleWi(wo, sampler.Get2D())

		if f.IsBlack() || pdf == 0.0 {
			break
		}

		beta = beta.Multiply(f).Scale(math.Abs(wi.Dot(isect.Normal)) / pdf)
		specularBounce = (bsdfType & BSDF_SPECULAR) != 0
		ray = isect.SpawnRay(wi)

		// Russian roulette
		if bounces > 3 {
			continueProbability := math.Min(0.95, beta.Y())
			if sampler.Get1D() > continueProbability {
				break
			}
			beta = beta.Scale(1.0 / continueProbability)
		}
	}
	return L
}

func NextEventEstimation(isect *Intersection, scene *Scene, sampler Sampler) *Color {
	if len(scene.Lights) == 0 {
		fmt.Println("Warning: No light exists!")
		return NewColor(0.0, 0.0, 0.0)
	}

	Ld := NewColor(0.0, 0.0, 0.0)
	lightPdf, bsdfPdf := 0.0, 0.0

	// Light sampling
	numLights := len(scene.Lights)
	lightID := int(sampler.Get1D() * Float(numLights))
	light := scene.Lights[lightID]

	Li, wi, lightPdf, vis := light.SampleLi(isect, sampler.Get2D())
	if !Li.IsBlack() && lightPdf > 0.0 && vis.Unoccluded(scene) {
		bsdfPdf = isect.Bsdf().Pdf(wi, isect.Wo)
		weight := powerHeuristic(lightPdf, bsdfPdf)
		f := isect.Bsdf().Eval(wi, isect.Wo).Scale(math.Abs(isect.Normal.Dot(wi)))
		L := f.Multiply(Li).Scale(weight / lightPdf)
		Ld = Ld.Add(L)
	}

	// BSDF sampling
	lightPdf, bsdfPdf = 0.0, 0.0
	f, wi, bsdfPdf, sampledType := isect.Bsdf().SampleWi(isect.Wo, sampler.Get2D())
	if !f.IsBlack() && bsdfPdf > 0.0 {
		sampledSpecular := (sampledType & BSDF_SPECULAR) != 0
		f = f.Scale(math.Abs(isect.Normal.Dot(wi)))
		weight := 1.0
		if !sampledSpecular {
			lightPdf = light.PdfLi(isect, wi)
			if lightPdf <= 0.0 {
				return Ld.Scale(Float(numLights))
			}
			weight = powerHeuristic(bsdfPdf, lightPdf)
		}

		var testIsect Intersection
		ray := NewRay(isect.Pos, wi)
		isIntersect := scene.Intersect(ray, &testIsect)
		Li := NewColor(0.0, 0.0, 0.0)
		if isIntersect {
			rr := testIsect.Le(ray.Dir.Negate())
			Li = Li.Add(rr)
		} else {
			for _, l := range scene.Lights {
				rr := l.LeWithRay(ray)
				Li = Li.Add(rr)
			}
		}

		L := f.Multiply(Li).Scale(weight / bsdfPdf)
		Ld = Ld.Add(L)
	}

	return Ld.Scale(Float(numLights))
}

func powerHeuristic(f, g Float) Float {
	return (f * f) / ((f * f) + (g * g))
}
