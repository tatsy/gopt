package integrator

import (
    "fmt"
    . "core"
    . "accelerator"
)

type PathIntegrator struct {
}

func (integrator *PathIntegrator) Render(bvh *Bvh, sensor Sensor, sampler Sampler) {
    width := sensor.Film().Width
    height := sensor.Film().Height
    film := sensor.Film()
    sem := make(chan Semaphore, width)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            go func(x, y int) {
                ray := sensor.SpawnRay(x, y)
                L := integrator.Li(bvh, ray, sampler)
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

func (integrator *PathIntegrator) Li(bvh *Bvh, ray *Ray, sampler Sampler) Color {
    var isect Intersection
    if bvh.Intersect(ray, &isect) {
        return Color{1.0, 1.0, 0.0}
    }
    return Color{0.0, 0.0, 0.0}
}
