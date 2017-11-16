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

	for bounces := 0; ; bounces++ {
		var isect Intersection
		isIntersect := scene.Intersect(ray, &isect)
		if bounces == 0 {
			if isIntersect {
				Ld := isect.Le(ray.Dir.Negate())
				L = L.Add(beta.Multiply(Ld))
			} else {
				for _, l := range scene.Lights {
					Ld := l.LeWithRay(ray)
					L = L.Add(beta.Multiply(Ld))
				}
			}
		}

		if !isIntersect || bounces >= maxBounces {
			break
		}

		Ld := NextEventEstimation(&isect, scene, sampler)
		L = L.Add(beta.Multiply(Ld))

		wo := ray.Dir.Negate()
		f, wi, pdf, _ := isect.Bsdf().SampleWi(wo, sampler.Get2D())

		if f.IsBlack() || pdf == 0.0 {
			break
		}

		beta = beta.Multiply(f).Scale(math.Abs(wi.Dot(isect.Normal)) / pdf)
		ray = isect.SpawnRay(wi)

		// Russian roulette
		if bounces > 3 {
			continueProbability := math.Min(0.95, beta.MaxComponent())
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

	var lightPdf, bsdfPdf Float
	Ld := NewColor(0.0, 0.0, 0.0)

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
	f, wi, bsdfPdf, bsdfType := isect.Bsdf().SampleWi(isect.Wo, sampler.Get2D())
	if !f.IsBlack() && bsdfPdf > 0.0 {
		f = f.Scale(math.Abs(isect.Normal.Dot(wi)))

		var lightIsect Intersection
		ray := isect.SpawnRay(wi)
		isIntersect := scene.Intersect(ray, &lightIsect)

		Li = NewColor(0.0, 0.0, 0.0)
		if isIntersect {
			if lightIsect.IsHitLight(light) {
				Li = lightIsect.Le(ray.Dir.Negate())
			}
		} else {
			Li = light.LeWithRay(ray)
		}

		if !Li.IsBlack() {
			lightPdf = light.PdfLi(isect, wi)
			weight := 1.0
			sampleSpecular := (bsdfType & BSDF_SPECULAR) != 0
			if sampleSpecular {
				weight = powerHeuristic(bsdfPdf, lightPdf)
			}
			L := f.Multiply(Li).Scale(weight / bsdfPdf)
			Ld = Ld.Add(L)
		}
	}

	return Ld.Scale(Float(numLights))
}

func powerHeuristic(f, g Float) Float {
	return (f * f) / ((f * f) + (g * g))
}
