package main

import (
    "fmt"
    . "core"
    . "accelerator"
    . "bsdf"
    . "shape"
    . "sensor"
    . "sampler"
    . "integrator"
)

const (
    fileName = "data/gopher.obj"
)

func main() {
    triMesh := TriMesh{}
    success := triMesh.Load(fileName)
    if !success {
        panic("Failed to load file!")
    }
    fmt.Printf("%d triangles\n", len(triMesh.Triangles))

    numTris := len(triMesh.Triangles)
    primitives := make([]Primitive, numTris)

    bsdf := LambertBsdf{Color{1.0, 1.0, 1.0}}
    for i := range triMesh.Triangles {
        primitives[i] = NewPrimitive(&triMesh.Triangles[i], &bsdf)
    }
    bvh := NewBvh(primitives)

    film := NewFilm(256, 256)
    sensor := NewPerspectiveSensor(
        Vector3d{5.0, 5.0, 5.0},  // Center
        Vector3d{0.0, 0.0, 0.0},  // To
        Vector3d{0.0, 1.0, 0.0},  // Up
        45.0,  // Fov
        film.Aspect(),  // Aspect
        0.1,  // Near clip
        1000.0,  // Far clip
        film,
    )

    sampler := &IndependentSampler{}
    integrator := PathIntegrator{}
    integrator.Render(bvh, sensor, sampler)
}
