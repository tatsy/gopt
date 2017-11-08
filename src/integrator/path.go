package integrator

import (
    . "../core"
)

type PathIntegrator struct {
}

func (integrator PathIntegrator) Render(bvh Bvh, sensor Sensor, sampler Sampler) {
    width := sensor.Film().Width
    height := sensor.Film().Height
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            ray := sensor.SpawnRay(x, y)
            L := integrator.Li(bvh, ray, sampler)
            sensor.Film().Update(x, y, L)
        }
    }
    sensor.Film().Save("image.jpg")
}

func (integrator PathIntegrator) Li(bvh Bvh, ray Ray, sampler Sampler) Color {
    var isect Intersection
    if bvh.Intersect(ray, &isect) {
        return Color{1.0, 1.0, 0.0}
    }
    return Color{0.0, 0.0, 0.0}
}
