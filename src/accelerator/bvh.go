package accelerator

import (
    "sort"
    . "core"
)

type BvhNode struct {
    left, right *BvhNode
    shape Shape
    bbox Bounds3d
}

func NewLeafNode(shape Shape) (node *BvhNode) {
    node = &BvhNode{}
    node.left = nil
    node.right = nil
    node.shape = shape
    node.bbox = shape.Bounds()
    return
}

func NewForkNode(left *BvhNode, right *BvhNode, b Bounds3d) (node *BvhNode) {
    node = &BvhNode{}
    node.left = left
    node.right = right
    node.shape = nil
    node.bbox = b
    return
}

func (node *BvhNode) IsLeaf() bool {
    return node.left == nil && node.right == nil
}

type SortItem struct {
    v Vector3d
    i int
}

type AxisSorter struct {
    Items []SortItem
    Axis int
}

func (a *AxisSorter) Len() int {
    return len(a.Items)
}

func (a *AxisSorter) Swap(i, j int) {
    a.Items[i], a.Items[j] = a.Items[j], a.Items[i]
}

func (a *AxisSorter) Less(i, j int) bool {
    v1 := a.Items[i].v
    v2 := a.Items[j].v
    return v1.NthElement(a.Axis) < v2.NthElement(a.Axis)
}

type Bvh struct {
    primitives []Primitive
    nodes []BvhNode
    root *BvhNode
}

func NewBvh(primitives []Primitive) (bvh Bvh) {
    bvh.primitives = primitives
    bvh.root = NewBvhSub(&bvh, primitives, 0)
    return
}

func NewBvhSub(bvh *Bvh, primitives []Primitive, axis int) *BvhNode {
    if len(primitives) == 1 {
        node := NewLeafNode(primitives[0].Shape)
        bvh.nodes = append(bvh.nodes, *node)
        return node
    }

    bbox := NewBounds3d()
    items := make([]SortItem, len(primitives))
    for i := range primitives {
        b := primitives[i].Shape.Bounds()
        bbox.Merge(b)
        items[i] = SortItem{b.Center(), i}
    }

    axisSorter := &AxisSorter{items, axis % 3}
    sort.Sort(axisSorter)

    newPrimitives := make([]Primitive, len(primitives))
    for i := range axisSorter.Items {
        newPrimitives[i] = primitives[axisSorter.Items[i].i]
    }

    iHalf := len(primitives) / 2
    leftNode := NewBvhSub(bvh, newPrimitives[:iHalf], (axis + 1) % 3)
    rightNode := NewBvhSub(bvh, newPrimitives[iHalf:], (axis + 1) % 3)

    node := NewForkNode(leftNode, rightNode, bbox)
    bvh.nodes = append(bvh.nodes, *node)
    return node
}

func (bvh *Bvh) Intersect(r Ray, isect *Intersection) bool {
    return IntersectSub(bvh.root, &r, isect)
}

func IntersectSub(node *BvhNode, r *Ray, isect *Intersection) bool {
    ret := false

    // Leaf node
    if node.IsLeaf() {
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
            ret = ret || IntersectSub(node.left, r, isect)
        }
        if node.right != nil {
            ret = ret || IntersectSub(node.right, r, isect)
        }
    }
    return ret
}
