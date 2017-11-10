package sensor

import (
    "math"
    . "core"
)

type PerspectiveSensor struct {
    center, unitU, unitV, unitW *Vector3d
    fov, aspect, nearClip, farClip Float
    film *Film
}

func NewPerspectiveSensor(
    center *Vector3d,
    target *Vector3d,
    up *Vector3d,
    fov Float,
    aspect Float,
    nearClip Float,
    farClip Float,
    film *Film) *PerspectiveSensor {
    sensor := &PerspectiveSensor{}
    sensor.center = center
    to := target.Subtract(center)
    sensor.unitU = to.Cross(up).Normalized()
    sensor.unitV = to.Cross(sensor.unitU).Normalized()
    sensor.unitW = to.Normalized()
    sensor.fov = fov
    sensor.aspect = aspect
    sensor.nearClip = nearClip
    sensor.farClip = farClip
    sensor.film = film
    return sensor
}

func (sensor *PerspectiveSensor) Film() *Film {
    return sensor.film
}

func (sensor *PerspectiveSensor) SpawnRay(x, y Float) *Ray {
    width := sensor.film.Width
    height := sensor.film.Height

    u := (x / Float(width)) - 0.5
    v := (y / Float(height)) - 0.5

    screenHeight := 2.0 * math.Tan(sensor.fov * 0.5) * sensor.nearClip
    screenWidth := sensor.aspect * screenHeight

    targetX := u * screenWidth
    targetY := v * screenHeight
    targetZ := sensor.nearClip
    direction := sensor.unitU.Scale(targetX).
                 Add(sensor.unitV.Scale(targetY)).
                 Add(sensor.unitW.Scale(targetZ)).
                 Normalized()
    return NewRay(sensor.center, direction)
}
