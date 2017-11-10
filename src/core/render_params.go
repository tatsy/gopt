package core

import (
    "fmt"
    "strconv"
)

type RenderParams struct {
    entries map[string]string
}

func NewRenderParams() *RenderParams {
    p := &RenderParams{}
    p.entries = make(map[string]string)
    return p
}

func (p *RenderParams) AddEntry(name string, value string) {
    p.entries[name] = value
}

func (p *RenderParams) GetString(name string) string {
    value, ok := p.entries[name]
    if !ok {
        panic(fmt.Sprintf("String name \"%s\" not found!", name))
    }
    return value
}

func (p *RenderParams) GetInt(name string) int {
    value, ok := p.entries[name]
    if !ok {
        panic(fmt.Sprintf("Int name \"%s\" not found!", name))
    }

    ret, err := strconv.Atoi(value)
    if err != nil {
        panic(err)
    }
    return ret
}

func (p *RenderParams) GetFloat(name string) Float {
    value, ok := p.entries[name]
    if !ok {
        panic(fmt.Sprintf("Float name \"%s\" not found!", name))
    }
    ret, err := strconv.ParseFloat(value, 64)
    if err != nil {
        panic(err)
    }
    return ret
}

func (p *RenderParams) GetVector3d(name string) *Vector3d {
    value, ok := p.entries[name]
    if !ok {
        panic(fmt.Sprintf("Vector3d name \"%s\" not found!", name))
    }

    var x, y, z Float
    _, err := fmt.Sscanf(value, "(%f, %f, %f)", &x, &y, &z)
    if err != nil {
        panic(err)
    }
    return NewVector3d(x, y, z)
}
