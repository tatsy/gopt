package accelerator

import (
    "sort"
    "math"
    "math/rand"
    "testing"
    . "core"
    . "shape"
    . "bsdf"
)

func TestAxisSorter(t *testing.T) {
    items := make([]*SortItem, 100)
    for i := range items {
        v := NewVector3d(
            rand.Float64(),
            rand.Float64(),
            rand.Float64(),
        )
        items[i] = &SortItem{v, i}
    }

    k := rand.Intn(3)
    a := &AxisSorter{items, k}
    sort.Sort(a)

    for i := 0; i < len(items) - 1; i++ {
        if items[i].v.NthElement(k) >= items[i + 1].v.NthElement(k) {
            t.Errorf("Item sorting failed!")
            break
        }
    }
}

func TestBvhIntersection(t *testing.T) {
    triMesh := NewTriMeshFromFile("../../data/cube.obj")
    prims := make([]*Primitive, triMesh.NumFaces())
    bsdf := &LambertBsdf{}
    for i := range triMesh.Triangles {
        prims[i] = NewPrimitive(triMesh.Triangles[i], bsdf)
    }
    bvh := NewBvh(prims)

    numTrials := 100
    for trial := 0; trial < numTrials; trial++ {
        org := NewVector3d(
            rand.Float64(),
            rand.Float64(),
            rand.Float64()).Scale(2.0)
        dir := NewVector3d(
            rand.Float64(),
            rand.Float64(),
            rand.Float64()).Scale(2.0)
        ray := NewRay(org, dir)

        var isect Intersection
        actual := bvh.Intersect(ray, &isect)
        actualDist := isect.HitDist

        expected := false
        expectedDist := Infinity
        for _, p := range bvh.primitives {
            var temp Intersection
            if p.Shape.Intersect(ray, &temp) {
                expected = true
                expectedDist = math.Min(expectedDist, temp.HitDist)
            }
        }

        if actual != expected {
            t.Errorf("Intersection test failed:\n%v\nexpected: %v\nactual: %v", ray, expected, actual)
        } else if actual && expected {
            if math.Abs(actualDist - expectedDist) >= 1.0e-8 {
                t.Errorf("Intersection distances differ: %f != %f", expectedDist, actualDist)
            }
        }
    }
}
