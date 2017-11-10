package integrator

import (
    "fmt"
    . "core"
)

type PathIntegrator struct {
}

func (integrator *PathIntegrator) Render(scene *Scene, sensor Sensor, sampler Sampler) {
    width := sensor.Film().Width
    height := sensor.Film().Height
    film := sensor.Film()
    sem := make(chan Semaphore, width)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            go func(x, y int) {
                ray := sensor.SpawnRay(x, y)
                L := integrator.Li(scene, ray, sampler)
                film.Update(x, y, L)
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

func (integrator *PathIntegrator) Li(scene *Scene, r *Ray, sampler Sampler) *Color {
    maxBounces := 16
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

        wo := ray.Dir.Negate()
        f, wi, pdf, bsdfType := isect.Bsdf().SampleWi(wo, sampler.Get2D())

        if f.IsBlack() || pdf == 0.0 {
            break
        }

        beta = beta.Multiply(f).Scale(wi.Dot(isect.Normal) / pdf)
        specularBounce = (bsdfType & BSDF_SPECULAR) != 0
        ray = isect.SpawnRay(wi)
    }

    if !L.IsBlack() {
        fmt.Println(L)
    }

    return L
}
