package core

import (
    "sort"
)

type BvhNode struct {
    left, right *BvhNode
    shape Shape
    bbox Bounds3d
}

func NewLeafNode(shape Shape) (node BvhNode) {
    node.left = nil
    node.right = nil
    node.shape = shape
    node.bbox = shape.Bounds()
    return
}

func NewForkNode(left *BvhNode, right *BvhNode, b Bounds3d) (node BvhNode) {
    node.left = left
    node.right = right
    node.shape = nil
    node.bbox = b
    return
}

type AxisSorter struct {
    Primitives []Primitive
    Axis int
}

func (a AxisSorter) Len() int {
    return len(a.Primitives)
}

func (a AxisSorter) Swap(i, j int) {
    a.Primitives[i], a.Primitives[j] = a.Primitives[j], a.Primitives[i]
}

func (a AxisSorter) Less(i, j int) bool {
    v1 := a.Primitives[i].Shape.Bounds().Center()
    v2 := a.Primitives[j].Shape.Bounds().Center()
    return v1.NthElement(a.Axis) < v2.NthElement(a.Axis)
}


type Bvh struct {
    primitives []Primitive
    nodes []BvhNode
}

func NewBvh(primitives []Primitive) (bvh Bvh) {
    bvh.primitives = primitives
    NewBvhSub(&bvh, primitives, 0)
    return
}

func NewBvhSub(bvh *Bvh, primitives []Primitive, axis int) *BvhNode {
    if len(primitives) == 1 {
        node := NewLeafNode(primitives[0].Shape)
        bvh.nodes = append(bvh.nodes, node)
        return &node
    }

    bbox := NewBounds3d()
    for i := range primitives {
        bbox.Merge(primitives[i].Shape.Bounds())
    }

    axisSorter := AxisSorter{primitives, axis % 3}
    sort.Sort(axisSorter)

    iHalf := len(primitives) / 2
    leftNode := NewBvhSub(bvh, primitives[:iHalf], (axis + 1) % 3)
    rightNode := NewBvhSub(bvh, primitives[iHalf:], (axis + 1) % 3)

    node := NewForkNode(leftNode, rightNode, bbox)
    bvh.nodes = append(bvh.nodes, node)
    return &node
}

func (bvh Bvh) Intersect(r Ray, isect *Intersection) bool {
    return IntersectSub(&bvh.nodes[0], &r, isect)
    // for i := range bvh.primitives {
    //     var isect Intersection
    //     if bvh.primitives[i].Shape.Intersect(r, &isect) {
    //         return true
    //     }
    // }
    // return false
}

func IntersectSub(node *BvhNode, r *Ray, isect *Intersection) bool {
    ret := false

    // Leaf node
    if node.left == nil && node.right == nil {
        var temp Intersection
        if node.shape.Intersect(*r, &temp) {
            ret = true
            *isect = temp
            r.MaxDist = temp.HitDist
        }
        return ret
    }

    // Fork node
    var tMin, tMax Float
    if node.bbox.Intersect(*r, &tMin, &tMax) {
        if node.left != nil {
            var temp Intersection
            if IntersectSub(node.left, r, &temp) {
                ret = true
                *isect = temp
                r.MaxDist = temp.HitDist
            }
        }

        if node.right != nil {
            var temp Intersection
            if IntersectSub(node.right, r, &temp) {
                ret = true
                *isect = temp
                r.MaxDist = temp.HitDist
            }
        }
    }
    return ret
}
