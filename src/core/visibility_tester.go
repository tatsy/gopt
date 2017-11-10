package core

type VisibilityTester struct {
    p0, p1 *Vector3d
}

func NewVisibilityTester(p0, p1 *Vector3d) *VisibilityTester {
    vt := &VisibilityTester{}
    vt.p0, vt.p1 = p0, p1
    return vt
}

func (vt *VisibilityTester) Unoccluded(scene *Scene) bool {
    ray := NewRayBetweenPoints(vt.p0, vt.p1)
    var isect Intersection
    return !scene.Intersect(ray, &isect)
}
