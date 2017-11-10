package core

import (
    "fmt"
)

type RenderParams struct {
    ints map[string]int
    floats map[string]Float
}

func NewRenderParams() *RenderParams {
    p := &RenderParams{}
    p.ints = make(map[string]int)
    p.floats = make(map[string]Float)
    return p
}

func (p *RenderParams) AddInt(name string, value int) {
    p.ints[name] = value
}

func (p *RenderParams) GetInt(name string) (int, error) {
    value, ok := p.ints[name]
    if !ok {
        e := fmt.Errorf("Int name \"%s\" not found!", name)
        return value, e
    }
    return value, nil
}

func (p *RenderParams) AddFloat(name string, value Float) {
    p.floats[name] = value
}

func (p *RenderParams) GetFloat(name string) (Float, error) {
    value, ok := p.floats[name]
    if !ok {
        e := fmt.Errorf("Float name \"%s\" not found!", name)
        return value, e
    }
    return value, nil
}
