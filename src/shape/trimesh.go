package shape

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    . "core"
)

type TriMesh struct {
    Triangles []*Triangle
}

func NewTriMeshFromFile(filename string) *TriMesh {
    triMesh := &TriMesh{}
    triMesh.Load(filename)
    return triMesh
}

func (triMesh *TriMesh) NumFaces() int {
    return len(triMesh.Triangles)
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

    // Read file content
    var positions []*Vector3d
    var normals []*Vector3d
    var triangles []*Triangle
    for scanner.Scan() {
        line := scanner.Text()

        // Check comment line
        isSkip := true
        for i := range line {
            if line[i] != ' ' {
                if line[i] == '#' {
                    break
                }
            } else {
                isSkip = false
            }
        }

        if isSkip {
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
            triangles = append(triangles, NewTriangle(
                [3]*Vector3d{positions[i - 1], positions[j - 1], positions[k - 1]},
                [3]*Vector3d{normals[ni - 1], normals[nj - 1], normals[nk - 1]},
            ))

        case "mtllib":
            continue
        }
    }

    triMesh.Triangles = triangles
    return true
}
