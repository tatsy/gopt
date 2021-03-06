package core

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestNewBounds3d(t *testing.T) {
	b := NewBounds3d()
	if b.MinPos.X != Infinity || b.MinPos.Y != Infinity || b.MinPos.Z != Infinity ||
		b.MaxPos.X != -Infinity || b.MaxPos.Y != -Infinity || b.MaxPos.Z != -Infinity {
		t.Error("Initialization failed")
	}
}

func TestMerge(t *testing.T) {
	b1 := NewBounds3dMinMax(
		NewVector3d(0.0, 0.0, 0.0),
		NewVector3d(1.0, 1.0, 1.0),
	)

	b2 := NewBounds3dMinMax(
		NewVector3d(0.5, 0.5, 0.5),
		NewVector3d(2.0, 2.0, 2.0),
	)

	b3 := b1
	b3.Merge(b2)
	if !b3.MinPos.Equals(b1.MinPos) && !b3.MaxPos.Equals(b2.MaxPos) {
		t.Error("Failed")
	}
}

func TestMergePoint(t *testing.T) {
	b := NewBounds3d()
	pMin := NewVector3d(Infinity, Infinity, Infinity)
	pMax := NewVector3d(-Infinity, -Infinity, -Infinity)
	for i := 0; i < 100; i++ {
		p := NewVector3d(rand.Float64(), rand.Float64(), rand.Float64())
		b = b.MergePoint(p)
		pMin = pMin.Minimum(p)
		pMax = pMax.Maximum(p)
	}

	if !pMin.Equals(b.MinPos) || !pMax.Equals(b.MaxPos) {
		t.Errorf("MergePoint failed: (%v, %v) != (%v, %v)", pMin, pMax, b.MinPos, b.MaxPos)
	}

}

func TestBounds3dIntersect(t *testing.T) {
	b1 := NewBounds3dMinMax(
		NewVector3d(0.0, 0.0, 0.0),
		NewVector3d(1.0, 1.0, 1.0),
	)
	r := NewRay(
		NewVector3d(1.5, 0.5, 0.5),
		NewVector3d(-1.0, 0.0, 0.0),
	)

	var tMin, tMax Float
	if !b1.Intersect(r, &tMin, &tMax) {
		t.Error("Failed")
	}

	if tMin != 0.5 || tMax != 1.5 {
		t.Errorf("Failed: %f vs %f", tMin, 0.5)
	}
}

func TestBounds3dCentroid(t *testing.T) {
	b := NewBounds3dMinMax(
		NewVector3d(1.0, 2.0, 3.0),
		NewVector3d(2.0, 3.0, 4.0),
	)
	c := b.Centroid()
	if c.X != 1.5 || c.Y != 2.5 || c.Z != 3.5 {
		t.Errorf("Unexpected centroid: %v", c)
	}
}

func TestMaxExtent(t *testing.T) {
	testCases := map[[6]Float]int{
		{-1.0, 2.0, 3.0, 4.0, 5.0, 6.0}: 0,
		{1.0, -2.0, 3.0, 4.0, 5.0, 6.0}: 1,
		{1.0, 2.0, -3.0, 4.0, 5.0, 6.0}: 2,
	}

	for v, n := range testCases {
		b := NewBounds3dMinMax(
			NewVector3d(v[0], v[1], v[2]),
			NewVector3d(v[3], v[4], v[5]),
		)

		if b.MaxExtent() != n {
			t.Errorf("Max extent of %v != %d", b, n)
		}
	}
}
