package integrator

import (
    "fmt"
    "math"
    . "core"
)

type PathIntegrator struct {
}

func (integrator *PathIntegrator) Render(scene *Scene, sensor Sensor, sampler Sampler, params *RenderParams) {
    width := sensor.Film().Width
    height := sensor.Film().Height
    numSamples := params.GetInt("integrator.num-samples")
    maxBounces := params.GetInt("integrator.max-bounces")
    film := sensor.Film()
    sem := make(chan Semaphore, width)

    for s := 0; s < numSamples; s++ {
        fmt.Printf("Sample (%d / %d):\n", s + 1, numSamples)
        for y := 0; y < height; y++ {
            for x := 0; x < width; x++ {
                go func(x, y int) {
                    subPos := sampler.Get2D()
                    px := Float(x) + subPos.X
                    py := Float(y) + subPos.Y
                    ray := sensor.SpawnRay(px, py)
                    L := integrator.Li(scene, ray, sampler, maxBounces)
                    film.AddSample(px, py, L)
                    sem <- Semaphore{}
                } (x, y)
            }

            for x := 0; x < width; x++ {
                <-sem
            }

            ProgressBar(y + 1, height)
        }
        fmt.Println()
        film.Save("image.jpg")
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

        beta = beta.Multiply(f).Scale(wi.Dot(isect.Normal) / pdf)
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

    numLights := len(scene.Lights)
    lightId := int(sampler.Get1D() * Float(numLights))
    light := scene.Lights[lightId]

    Li, wi, pdf, vis := light.SampleLi(isect, sampler.Get2D())
    if pdf == 0.0 {
        return NewColor(0.0, 0.0, 0.0)
    }

    if !vis.Unoccluded(scene) {
        return NewColor(0.0, 0.0, 0.0)
    }

    f := isect.Bsdf().Eval(wi, isect.Wo).Scale(math.Abs(isect.Normal.Dot(wi)))
    return f.Multiply(Li).Scale(1.0 / pdf).Scale(Float(numLights))
}
