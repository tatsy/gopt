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
    "os"
    "log"
    "os/signal"
    "runtime/pprof"
)

const (
    fileName = "data/gopher.obj"
)

func main() {
    cpuprofile := "profile.prof"
    f, err := os.Create(cpuprofile)
    if err != nil {
       log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
       for sig := range c {
           log.Printf("captured %v, stopping profiler and exiting...", sig)
           pprof.StopCPUProfile()
           os.Exit(1)
        }
    }()


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
