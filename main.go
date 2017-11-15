package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/tatsy/gopt/src/accelerator"
	. "github.com/tatsy/gopt/src/core"
	. "github.com/tatsy/gopt/src/integrator"
	. "github.com/tatsy/gopt/src/light"
	. "github.com/tatsy/gopt/src/sampler"
	. "github.com/tatsy/gopt/src/sensor"
	. "github.com/tatsy/gopt/src/shape"
)

var objFiles = []string{
	"data/gopher.obj",
	"data/floor.obj",
}

var lightFiles = []string{
	"data/sphere.obj",
}

type JsonObject struct {
	Name   string      `json:"name"`
	Params []JsonParam `json:"parameters"`
}

type JsonParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	// cpuprofile := "profile.prof"
	// f, err := os.Create(cpuprofile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// go func() {
	// 	for sig := range c {
	// 		log.Printf("captured %v, stopping profiler and exiting...", sig)
	// 		pprof.StopCPUProfile()
	// 		os.Exit(1)
	// 	}
	// }()

	// Parse command line args
	var jsonFile string
	var numThreads int
	flag.StringVar(&jsonFile, "input", "", "Input JSON file")
	flag.StringVar(&jsonFile, "i", "", "Input JSON file")
	flag.IntVar(&numThreads, "nthreads", -1, "Number of threads to use")
	flag.Parse()
	if jsonFile == "" {
		flag.Usage()
		os.Exit(0)
	}

	if numThreads > 0 {
		runtime.GOMAXPROCS(numThreads)
	}

	// File info
	absPath, _ := filepath.Abs(jsonFile)
	absDir, fileName := filepath.Split(absPath)
	baseName := strings.TrimRight(fileName, filepath.Ext(fileName))
	outFile := filepath.Join(absDir, baseName+".jpg")

	// Parse JSON file
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var jsonObjects []JsonObject
	json.Unmarshal(bytes, &jsonObjects)

	// Parse objects, lights and parameters
	params := NewRenderParams()
	params.AddEntry("outfile", outFile)
	var primitives []*Primitive
	var lights []Light
	for _, obj := range jsonObjects {
		switch obj.Name {
		case "shape":
			for _, par := range obj.Params {
				if par.Name == "obj" {
					fileName := filepath.Join(absDir, par.Value)
					triMesh := TriMesh{}
					if !triMesh.Load(fileName) {
						panic("Failed to load file!")
					}

					for _, p := range triMesh.Primitives {
						primitives = append(primitives, p)
					}
				}
			}
		case "light":
			meshes := make([]*Primitive, 0)
			isDone := false
			for _, par := range obj.Params {
				if par.Name == "obj" {
					fileName := filepath.Join(absDir, par.Value)
					triMesh := TriMesh{}
					if !triMesh.Load(fileName) {
						panic("Failed to load light file!")
					}
					meshes = append(meshes, triMesh.Primitives...)
				}
			}

			for _, par := range obj.Params {
				if par.Name == "radiance" {
					if isDone {
						panic("Multiple \"radiance\" is specified to single light!")
					}

					Le := NewColorWithString(par.Value)
					for _, p := range meshes {
						area := NewAreaLight(p.Shape(), Le)
						p.SetLight(area)
						lights = append(lights, area)
						primitives = append(primitives, p)
					}
					isDone = true
				}
			}
		default:
			for _, par := range obj.Params {
				name := obj.Name + "." + par.Name
				value := par.Value
				params.AddEntry(name, value)
			}
		}
	}

	// Create scene
	bvh := NewBvh(primitives)
	scene := NewScene(bvh, lights)

	imageWidth := params.GetInt("image.width")
	imageHeight := params.GetInt("image.height")
	film := NewFilm(imageWidth, imageHeight)
	sensor := NewSensor(params, film)

	sampler := NewIndependentSampler()
	integrator := PathIntegrator{}

	timer := NewTimer()
	timer.Start()
	defer timer.Stop()
	integrator.Render(scene, sensor, sampler, params)
}
