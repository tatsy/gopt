package core

type Point2d struct {
    X, Y Float
}

func (p *Point2d) NthElement(i int) Float {
    switch i {
    case 0:
        return p.X
    case 1:
        return p.Y
    default:
        panic("Element index out of range!")
    }
}
