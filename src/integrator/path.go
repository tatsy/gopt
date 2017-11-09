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

func (integrator *PathIntegrator) Li(scene *Scene, ray *Ray, sampler Sampler) *Color {
    var isect Intersection
    if scene.Intersect(ray, &isect) {
        return NewColor(1.0, 1.0, 0.0)
    }
    return NewColor(0.0, 0.0, 0.0)

    // maxBounces := 16
    // L := NewColor(0.0, 0.0, 0.0)
    // beta := NewColor(1.0, 1.0, 1.0)
    // specularBounce := false
    //
    // for bounce := 0; bounce < maxBounces; bounce++ {
    //     var isect Intersection
    //     isIntersect := scene.Intersect(ray, &isect)
    //     if !isIntersect {
    //         if bounce == 0 || specularBounce {
    //             //return scene.Lights
    //         } else {
    //             for _, l := range scene.Lights {
    //                 L = L.Add(l.Le(ray).Multiply(beta))
    //             }
    //             return L
    //         }
    //     }
    // }
    //
    // return L
}
