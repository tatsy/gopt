package core

import (
    "os"
    "log"
    "math"
    "image"
    "image/color"
    "image/jpeg"
)

type Film struct {
    Width, Height int
    data []*Color
    weights []Float
}

func NewFilm(width, height int) *Film {
    film := &Film{}
    film.Width = width
    film.Height = height
    film.data = make([]*Color, width * height)
    film.weights = make([]Float, width * height)
    for i := 0; i < width * height; i++ {
        film.data[i] = NewColor(0.0, 0.0, 0.0)
        film.weights[i] = 0.0
    }
    return film
}

func (film *Film) Aspect() Float {
    return Float(film.Width) / Float(film.Height)
}

func (film *Film) AddSample(x, y Float, color *Color) {
    dx := x - 0.5
    dy := y - 0.5
    weight := math.Min(0.5 - dx, 0.5 - dy)

    index := int(y) * film.Width + int(x)
    film.data[index] = film.data[index].Add(color.Scale(weight))
    film.weights[index] += weight
}

func (film *Film) Save(filename string) {
    file, err := os.Create(filename)
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }

    width := film.Width
    height := film.Height
    image := image.NewRGBA(image.Rect(0, 0, film.Width, film.Height))
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            index := y * width + x
            r := toInt(film.data[index].R / film.weights[index])
            g := toInt(film.data[index].G / film.weights[index])
            b := toInt(film.data[index].B / film.weights[index])
            color := color.RGBA{r, g, b, 255}
            image.Set(x, y, color)
        }
    }

    options := &jpeg.Options{Quality: 100}
    jpeg.Encode(file, image, options)
}

func toInt(x Float) uint8 {
    var v Float
    v = math.Max(0.0, math.Min(x, 1.0))
    v = math.Pow(v, 1.0 / 2.2)
    return uint8(v * 255.0)
}
