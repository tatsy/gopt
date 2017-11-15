package sensor

import (
	. "github.com/tatsy/gopt/src/core"
)

type PerspectiveSensor struct {
	cameraToWorld               *Transform
	focusDistance, apertureSize Float
	cameraToScreen              *Transform
	screenToCamera              *Transform
	film                        *Film
}

func NewPerspectiveSensor(
	cameraToWorld *Transform,
	focusDistance Float,
	apertureSize Float,
	fov Float,
	nearClip Float,
	farClip Float,
	film *Film) *PerspectiveSensor {
	sensor := new(PerspectiveSensor)
	sensor.cameraToWorld = cameraToWorld
	sensor.focusDistance = focusDistance
	sensor.apertureSize = apertureSize
	sensor.cameraToScreen = NewPerspective(fov, film.Aspect(), nearClip, farClip)
	sensor.screenToCamera = sensor.cameraToScreen.Inverted()
	sensor.film = film
	return sensor
}

func (sensor *PerspectiveSensor) Film() *Film {
	return sensor.film
}

func (sensor *PerspectiveSensor) SpawnRay(x, y Float, u *Vector2d) *Ray {
	width := sensor.film.Width
	height := sensor.film.Height
	screenX := 2.0*(x/Float(width)) - 1.0
	screenY := -2.0*(y/Float(height)) + 1.0
	dirCameraSpace := sensor.screenToCamera.
		ApplyToP(NewVector3d(screenX, screenY, 0.0)).
		Normalized()
	orgCameraSpace := NewVector3d(0.0, 0.0, 0.0)
	if sensor.apertureSize > 0.0 {
		pLens := SampleConcentricDisk(u).Scale(sensor.apertureSize)
		ft := sensor.focusDistance / dirCameraSpace.Z
		pFocus := orgCameraSpace.Add(dirCameraSpace.Scale(ft))

		orgCameraSpace = NewVector3d(pLens.X, pLens.Y, 0.0)
		dirCameraSpace = pFocus.Subtract(orgCameraSpace).Normalized()
	}
	dir := sensor.cameraToWorld.ApplyToV(dirCameraSpace)
	org := sensor.cameraToWorld.ApplyToP(orgCameraSpace)
	return NewRay(org, dir)
}
