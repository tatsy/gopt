package main

import (
    . "core"
    . "accelerator"
    . "shape"
    . "light"
    . "sensor"
    . "sampler"
    . "integrator"
    // "os"
    // "log"
    // "os/signal"
    // "runtime/pprof"
)

var objFiles = []string{
    "data/gopher.obj",
    "data/floor.obj",
}

var lightFiles = []string{
    "data/sphere.obj",
}

func main() {
    // cpuprofile := "profile.prof"
    // f, err := os.Create(cpuprofile)
    // if err != nil {
    //    log.Fatal(err)
    // }
    // pprof.StartCPUProfile(f)
    // defer pprof.StopCPUProfile()
    // c := make(chan os.Signal, 1)
    // signal.Notify(c, os.Interrupt)
    // go func() {
    //    for sig := range c {
    //        log.Printf("captured %v, stopping profiler and exiting...", sig)
    //        pprof.StopCPUProfile()
    //        os.Exit(1)
    //     }
    // }()

    params := NewRenderParams()
    params.AddInt("image.width", 640)
    params.AddInt("image.height", 360)
    params.AddInt("integrator.max-bounces", 16)
    params.AddInt("integrator.num-samples", 16)

    var primitives []*Primitive
    var lights []Light

    // Scene
    for _, fileName := range objFiles {
        triMesh := TriMesh{}
        if !triMesh.Load(fileName) {
            panic("Failed to load file!")
        }

        for _, p := range triMesh.Primitives {
            primitives = append(primitives, p)
        }
    }

    // Light
    for _, fileName := range lightFiles {
        lightMesh := TriMesh{}
        if !lightMesh.Load(fileName) {
            panic("Failed to load light file!")
        }

        Le := NewColor(8.0, 8.0, 8.0)
        for _, p := range lightMesh.Primitives {
            area := NewAreaLight(p.Shape(), Le)
            p.SetLight(area)
            lights = append(lights, area)
            primitives = append(primitives, p)
        }
    }

    bvh := NewBvh(primitives)
    scene := NewScene(bvh, lights)

    imageWidth, _ := params.GetInt("image.width")
    imageHeight, _ := params.GetInt("image.height")
    film := NewFilm(imageWidth, imageHeight)
    sensor := NewPerspectiveSensor(
        NewVector3d(-2.0, 4.0, 5.0),  // Center
        NewVector3d(0.0, 2.0, 0.0),  // To
        NewVector3d(0.0, 1.0, 0.0),  // Up
        45.0,  // Fov
        film.Aspect(),  // Aspect
        0.1,  // Near clip
        1000.0,  // Far clip
        film,
    )

    sampler := &IndependentSampler{}
    integrator := PathIntegrator{}

    timer := NewTimer()
    timer.Start()
    defer timer.Stop()
    integrator.Render(scene, sensor, sampler, params)
}
