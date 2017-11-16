package shape

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/tatsy/gopt/src/bsdf"
	. "github.com/tatsy/gopt/src/core"
)

type TriMesh struct {
	Primitives []*Primitive
}

func NewTriMeshFromFile(filename string) *TriMesh {
	triMesh := &TriMesh{}
	triMesh.Load(filename)
	return triMesh
}

func (triMesh *TriMesh) NumFaces() int {
	return len(triMesh.Primitives)
}

func (triMesh *TriMesh) Load(filename string) bool {
	// Open file
	fmt.Printf("Load: %s\n", filename)
	handler, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	defer handler.Close()
	if err != nil {
		panic(err)
	}

	// Create buffered reader
	scanner := bufio.NewScanner(handler)

	// Default material
	var defaultMat, currentMat Bxdf
	defaultMat = NewLambertReflection(NewColor(0.0, 0.0, 0.0))
	currentMat = defaultMat
	materials := make(map[string]Bxdf)

	// Read file content
	var positions []*Vector3d
	var normals []*Vector3d
	var texCoords []*Vector2d
	var primitives []*Primitive
	for scanner.Scan() {
		line := scanner.Text()

		// Check comment line
		if IsIgnorableLine(line) {
			continue
		}

		items := strings.Split(line, " ")
		switch items[0] {
		case "v":
			x, _ := strconv.ParseFloat(items[1], 64)
			y, _ := strconv.ParseFloat(items[2], 64)
			z, _ := strconv.ParseFloat(items[3], 64)
			positions = append(positions, NewVector3d(x, y, z))

		case "vn":
			nx, _ := strconv.ParseFloat(items[1], 64)
			ny, _ := strconv.ParseFloat(items[2], 64)
			nz, _ := strconv.ParseFloat(items[3], 64)
			normals = append(normals, NewVector3d(nx, ny, nz))

		case "vt":
			tx, _ := strconv.ParseFloat(items[1], 64)
			ty, _ := strconv.ParseFloat(items[2], 64)
			texCoords = append(texCoords, NewVector2d(tx, ty))

		case "f":
			var i, ti, ni, j, tj, nj, k, tk, nk int
			if len(items) >= 5 {
				panic(fmt.Sprintf("Only triangle mesh is supported: %v", items))
			}

			ParseFace(items[1], &i, &ti, &ni)
			ParseFace(items[2], &j, &tj, &nj)
			ParseFace(items[3], &k, &tk, &nk)

			var triangle *Triangle
			if i >= 1 && j >= 1 && k >= 1 {
				if ti >= 1 && tj >= 1 && tk >= 1 {
					if ni >= 1 && nj >= 1 && nk >= 1 {
						triangle = NewTriangleWithPTN(
							[3]*Vector3d{positions[i-1], positions[j-1], positions[k-1]},
							[3]*Vector2d{texCoords[ti-1], texCoords[tj-1], texCoords[tk-1]},
							[3]*Vector3d{normals[ni-1], normals[nj-1], normals[nk-1]},
						)
					} else {
						triangle = NewTriangleWithPT(
							[3]*Vector3d{positions[i-1], positions[j-1], positions[k-1]},
							[3]*Vector2d{texCoords[ti-1], texCoords[tj-1], texCoords[tk-1]},
						)
					}
				} else if ni >= 1 && nj >= 1 && nk >= 1 {
					triangle = NewTriangleWithPN(
						[3]*Vector3d{positions[i-1], positions[j-1], positions[k-1]},
						[3]*Vector3d{normals[ni-1], normals[nj-1], normals[nk-1]},
					)
				} else {
					triangle = NewTriangleWithP(
						[3]*Vector3d{positions[i-1], positions[j-1], positions[k-1]},
					)
				}
			}

			if triangle == nil {
				panic("Failed to parse triangle!")
			}

			if currentMat == nil {
				panic("Material is nil")
			}
			primitives = append(primitives, NewPrimitive(triangle, currentMat))

		case "usemtl":
			mat, isFound := materials[items[1]]
			if mat == nil || !isFound {
				panic(fmt.Sprintf("Material not found: %s", items[1]))
			}
			currentMat = mat

		case "mtllib":
			absName, _ := filepath.Abs(filename)
			currentDir := filepath.Dir(absName)
			materials = LoadMaterial(filepath.Join(currentDir, items[1]))
		}
	}

	triMesh.Primitives = primitives
	return true
}

func ParseFace(item string, i, ti, ni *int) {
	var n int
	*i, *ti, *ni = -1, -1, -1
	n, _ = fmt.Sscan(item, "%d/%d/%d", i, ti, ni)
	if n != 3 {
		n, _ = fmt.Sscanf(item, "%d/%d", i, ti)
		if n != 2 {
			n, _ = fmt.Sscanf(item, "%d//%d", i, ni)
			if n != 2 {
				n, _ = fmt.Sscanf(item, "%d", i)
				if n != 1 {
					panic(fmt.Sprintf("Face format is invalid: %s", item))
				}
			}
		}
	}
}

func IsIgnorableLine(line string) bool {
	for i := range line {
		if line[i] != ' ' {
			if line[i] == '#' {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func LoadMaterial(filename string) map[string]Bxdf {
	// Open file
	handler, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	defer handler.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Create buffered reader
	scanner := bufio.NewScanner(handler)

	materials := make(map[string]Bxdf)
	currentName := ""
	Kd := NewColor(0.5, 0.5, 0.5)
	Ks := NewColor(0.0, 0.0, 0.0)
	eta := 0.0
	for scanner.Scan() {
		line := scanner.Text()
		if IsIgnorableLine(line) {
			continue
		}

		items := strings.Split(line, " ")
		switch {
		case items[0] == "newmtl":
			if currentName != "" {
				materials[currentName] = NewBxdf(Kd, Ks, eta)
			}
			currentName = items[1]
			Kd = NewColor(0.5, 0.5, 0.5)
			Ks = NewColor(0.0, 0.0, 0.0)
			eta = 0.0

		case items[0] == "Kd":
			r, _ := strconv.ParseFloat(items[1], 64)
			g, _ := strconv.ParseFloat(items[2], 64)
			b, _ := strconv.ParseFloat(items[3], 64)
			Kd = NewColor(r, g, b)

		case items[0] == "Ks":
			r, _ := strconv.ParseFloat(items[1], 64)
			g, _ := strconv.ParseFloat(items[2], 64)
			b, _ := strconv.ParseFloat(items[3], 64)
			Ks = NewColor(r, g, b)

		case items[0] == "Ni":
			eta, _ = strconv.ParseFloat(items[1], 64)
		}
	}

	if currentName != "" {
		materials[currentName] = NewBxdf(Kd, Ks, eta)
	}

	return materials
}

func NewBxdf(Kd, Ks *Color, eta Float) Bxdf {
	switch {
	case eta != 0.0:
		return NewSpecularFresnel(Ks, Ks, 1.0, eta)
	case !Kd.IsBlack() && !Ks.IsBlack():
		return NewLambertReflection(Kd)
	case !Kd.IsBlack():
		return NewLambertReflection(Kd)
	case !Ks.IsBlack():
		return NewSpecularReflection(Ks)
	default:
		return NewLambertReflection(Kd)
	}
}
