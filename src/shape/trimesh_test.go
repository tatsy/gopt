package shape

import (
	"testing"
)

func TestTriMeshLoad(t *testing.T) {
	triMesh := NewTriMeshFromFile("../../testdata/cube.obj")
	if triMesh.NumFaces() != 12 {
		t.Errorf("# of cube mesh must be %d, detected %d", 12, triMesh.NumFaces())
	}

	defer func() {
		if p := recover(); p != nil {
		}
	}()
	NewTriMeshFromFile("invalid.name.obj")
	t.Errorf("Invalid file loading did not panic!")
}
