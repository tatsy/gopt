package shape

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "path/filepath"
    . "core"
    . "bsdf"
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
    handler, err := os.OpenFile(filename, os.O_RDONLY, 0600)
    defer handler.Close()
    if err != nil {
        log.Fatal(err)
    }

    // Create buffered reader
    scanner := bufio.NewScanner(handler)

    // Default material
    var defaultMat, currentMat Bxdf
    defaultMat = NewLambertReflection(NewColor(0.5, 0.5, 0.5))
    currentMat = defaultMat
    materials := make(map[string]Bxdf)

    // Read file content
    var positions []*Vector3d
    var normals []*Vector3d
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
            continue

        case "f":
            var i, ni, j, nj, k, nk int
            fmt.Sscanf(items[1], "%d//%d", &i, &ni)
            fmt.Sscanf(items[2], "%d//%d", &j, &nj)
            fmt.Sscanf(items[3], "%d//%d", &k, &nk)
            triangle := NewTriangle(
                [3]*Vector3d{positions[i - 1], positions[j - 1], positions[k - 1]},
                [3]*Vector3d{normals[ni - 1], normals[nj - 1], normals[nk - 1]},
            )
            primitives = append(primitives, NewPrimitive(triangle, currentMat))

        case "usemtl":
            mat, isFound := materials[items[1]]
            if isFound {
                currentMat = mat
            }

        case "mtllib":
            absName, _ := filepath.Abs(filename)
            currentDir := filepath.Dir(absName)
            materials = LoadMaterial(filepath.Join(currentDir, items[1]))
        }
    }

    triMesh.Primitives = primitives
    return true
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
    for scanner.Scan() {
        line := scanner.Text()
        if IsIgnorableLine(line) {
            continue
        }

        items := strings.Split(line, " ")
        switch {
        case items[0] == "newmtl":
            if currentName != "" {
                materials[currentName] = NewLambertReflection(Kd)
            }
            currentName = items[1]

        case items[0] == "Kd":
            r, _ := strconv.ParseFloat(items[1], 64)
            g, _ := strconv.ParseFloat(items[2], 64)
            b, _ := strconv.ParseFloat(items[3], 64)
            Kd = NewColor(r, g, b)
        }
    }

    return materials
}
