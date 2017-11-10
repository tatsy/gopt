package main

import (
    . "core"
    . "accelerator"
    . "bsdf"
    . "shape"
    . "light"
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
    lightFileName = "data/sphere.obj"
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

    var primitives []*Primitive
    var lights []Light

    // Scene
    triMesh := TriMesh{}
    if !triMesh.Load(fileName) {
        panic("Failed to load file!")
    }

    bsdf := NewLambertReflection(NewColor(1.0, 1.0, 1.0))
    for i := range triMesh.Triangles {
        primitives = append(primitives, NewPrimitive(triMesh.Triangles[i], bsdf))
    }

    // Light
    lightMesh := TriMesh{}
    if !lightMesh.Load(lightFileName) {
        panic("Failed to load light file!")
    }

    lightBsdf := NewLambertReflection(NewColor(0.0, 0.0, 0.0))
    Le := NewColor(4.0, 4.0, 4.0)
    for i := range lightMesh.Triangles {
        area := NewAreaLight(triMesh.Triangles[i], Le)
        lights = append(lights, area)
        primitives = append(primitives, NewLightPrimitive(triMesh.Triangles[i], lightBsdf, area))
    }

    bvh := NewBvh(primitives)
    scene := NewScene(bvh, lights)

    film := NewFilm(256, 256)
    sensor := NewPerspectiveSensor(
        NewVector3d(5.0, 5.0, 5.0),  // Center
        NewVector3d(0.0, 0.0, 0.0),  // To
        NewVector3d(0.0, 1.0, 0.0),  // Up
        45.0,  // Fov
        film.Aspect(),  // Aspect
        0.1,  // Near clip
        1000.0,  // Far clip
        film,
    )

    sampler := &IndependentSampler{}
    integrator := PathIntegrator{}
    integrator.Render(scene, sensor, sampler)
}
