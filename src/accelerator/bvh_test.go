package gopt

import (
	"math"
	"math/rand"
	"testing"

	. "github.com/tatsy/gopt/src/core"
	. "github.com/tatsy/gopt/src/shape"
)

type BvhTestSotable struct {
	values []int
}

func NewBvhTestSorable(values []int) *BvhTestSotable {
	s := &BvhTestSotable{}
	s.values = values
	return s
}

func (s *BvhTestSotable) Len() int {
	return len(s.values)
}

func (s *BvhTestSotable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s *BvhTestSotable) Less(i, j int) bool {
	return s.values[i] < s.values[j]
}

func (s *BvhTestSotable) Check(i int) bool {
	return s.values[i]%2 == 0
}

func TestNthElement(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		numElems := 100
		values := make([]int, numElems)
		for i := 0; i < numElems; i++ {
			values[i] = rand.Intn(numElems)
		}
		sortable := NewBvhTestSorable(values)
		splitPos := 40 //rand.Intn(numElems)
		SliceNthElement(sortable, 0, numElems, splitPos)

		success := true
		for i := 0; i < splitPos; i++ {
			for j := splitPos; j < numElems; j++ {
				if values[i] > values[j] {
					t.Errorf("Condition violated: %d < %d", values[i], values[j])
					success = false
				}
			}
		}

		if !success {
			t.Errorf("%v", values)
		}
	}
}

func TestPartition(t *testing.T) {
	for trial := 0; trial < TestTrials; trial++ {
		numElems := 100
		values := make([]int, numElems)
		for i := 0; i < numElems; i++ {
			values[i] = rand.Intn(numElems)
		}
		sortable := NewBvhTestSorable(values)

		splitPos := SlicePartition(sortable)
		for i := 0; i < splitPos; i++ {
			if !sortable.Check(i) {
				t.Errorf("Condition violated: %d %% 2 == 0", values[i])
			}
		}
		for i := splitPos; i < numElems; i++ {
			if sortable.Check(i) {
				t.Errorf("Condition violated: %d %% 2 != 0", values[i])
			}
		}
	}
}

func TestBvhIntersection(t *testing.T) {
	triMesh := NewTriMeshFromFile("../../scenes/gopher/gopher.obj")
	bvh := NewBvh(triMesh.Primitives)

	for trial := 0; trial < TestTrials; trial++ {
		org := NewVector3d(
			rand.Float64(),
			rand.Float64(),
			rand.Float64()).Scale(2.0)
		dir := NewVector3d(
			rand.Float64(),
			rand.Float64(),
			rand.Float64()).Scale(2.0)
		r1 := NewRay(org, dir)

		var isect Intersection
		actual := bvh.Intersect(r1, &isect)
		actualDist := Infinity
		if actual {
			actualDist = r1.Org.Subtract(isect.Pos).Length()
		}

		r2 := NewRay(org, dir)
		expected := false
		for _, p := range bvh.primitives {
			var temp Intersection
			if p.Intersect(r2, &temp) {
				expected = true
				isect = temp
			}
		}
		expectedDist := Infinity
		if expected {
			expectedDist = r2.Org.Subtract(isect.Pos).Length()
		}

		if actual != expected {
			t.Errorf("Intersection test failed:\n%v\nexpected: %v\nactual: %v", r1, expected, actual)
		} else if actual && expected {
			if math.Abs(actualDist-expectedDist) >= 1.0e-8 {
				t.Errorf("Intersection distances differ: %f != %f", expectedDist, actualDist)
			}
		}
	}
}
