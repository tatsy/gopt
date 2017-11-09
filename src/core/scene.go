package core

type Scene struct {
    bvh Accelerator
    Lights []Light
}

func NewScene(bvh Accelerator, lights []Light) *Scene {
    s := &Scene{}
    s.bvh = bvh
    s.Lights = lights
    return s
}

func (s *Scene) Intersect(ray *Ray, isect *Intersection) bool {
    return s.bvh.Intersect(ray, isect)
}
