package sensor

import (
	"fmt"

	. "github.com/tatsy/gopt/src/core"
)

func NewSensor(params *RenderParams, film *Film) Sensor {
	typeStr := params.GetString("sensor.type")
	switch typeStr {
	case "perspective":
		return NewPerspectiveSensor(
			params.GetVector3d("sensor.center"),      // Center
			params.GetVector3d("sensor.target"),      // Target
			params.GetVector3d("sensor.up"),          // Up
			params.GetFloat("sensor.fov"),            // Fov
			film.Aspect(),                            // Aspect
			params.GetFloat("sensor.focus-distance"), // Focus distance
			params.GetFloat("sensor.near-clip"),      // Near clip
			params.GetFloat("sensor.far-clip"),       // Far clip
			film,
		)
	default:
		panic(fmt.Sprintf("Unknown sensor type: %s", typeStr))
	}
}
