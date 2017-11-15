package sensor

import (
	"fmt"

	. "github.com/tatsy/gopt/src/core"
)

func NewSensor(params *RenderParams, film *Film) Sensor {
	typeStr := params.GetString("sensor.type")
	cameraToWorld := NewLookAt(
		params.GetVector3d("sensor.center"),
		params.GetVector3d("sensor.target"),
		params.GetVector3d("sensor.up"),
	)
	switch typeStr {
	case "perspective":
		return NewPerspectiveSensor(
			cameraToWorld,
			params.GetFloat("sensor.focus-distance"), // Focus distance
			params.GetFloat("sensor.aperture-size"),  // Aperture size
			params.GetFloat("sensor.fov"),            // Fov
			params.GetFloat("sensor.near-clip"),      // Near clip
			params.GetFloat("sensor.far-clip"),       // Far clip
			film,
		)
	default:
		panic(fmt.Sprintf("Unknown sensor type: %s", typeStr))
	}
}
