package accelerator

import (
    "sort"
    . "core"
)

type BvhNode struct {
    left, right *BvhNode
    shape Shape
    bbox *Bounds3d
}

func NewLeafNode(shape Shape) *BvhNode {
    node := &BvhNode{}
    node.left = nil
    node.right = nil
    node.shape = shape
    node.bbox = shape.Bounds()
    return node
}

func NewForkNode(left *BvhNode, right *BvhNode, b *Bounds3d) *BvhNode {
    node := &BvhNode{}
    node.left = left
    node.right = right
    node.shape = nil
    node.bbox = b
    return node
}

func (node *BvhNode) IsLeaf() bool {
    return node.left == nil && node.right == nil
}

type SortItem struct {
    v *Vector3d
    i int
}

type AxisSorter struct {
    Items []*SortItem
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
    primitives []*Primitive
    nodes []*BvhNode
    root *BvhNode
}

func NewBvh(primitives []*Primitive) *Bvh {
    bvh := &Bvh{}
    bvh.primitives = primitives
    bvh.root = NewBvhSub(bvh, primitives)
    return bvh
}

func NewBvhSub(bvh *Bvh, primitives []*Primitive) *BvhNode {
    if len(primitives) == 1 {
        node := NewLeafNode(primitives[0].Shape)
        bvh.nodes = append(bvh.nodes, node)
        return node
    }

    bbox := NewBounds3d()
    items := make([]*SortItem, len(primitives))
    for i := range primitives {
        b := primitives[i].Shape.Bounds()
        bbox.Merge(b)
        items[i] = &SortItem{b.Center(), i}
    }
    axis := bbox.MaxExtent()

    axisSorter := &AxisSorter{items, axis}
    sort.Sort(axisSorter)

    newPrimitives := make([]*Primitive, len(primitives))
    for i := range items {
        newPrimitives[i] = primitives[items[i].i]
    }

    iHalf := len(newPrimitives) / 2
    leftNode := NewBvhSub(bvh, newPrimitives[:iHalf])
    rightNode := NewBvhSub(bvh, newPrimitives[iHalf:])

    node := NewForkNode(leftNode, rightNode, bbox)
    bvh.nodes = append(bvh.nodes, node)
    return node
}

func (bvh *Bvh) Intersect(r *Ray, isect *Intersection) bool {
    stack := make([]*BvhNode, 40)
    pos := 0
    stack[pos] = bvh.root

    ret := false
    for pos >= 0 {
        // Pop stack
        node := stack[pos]
        pos -= 1

        if node.IsLeaf() {
            // Leaf node
            var temp Intersection
            if node.shape.Intersect(r, &temp) {
                ret = true
                *isect = temp
                r.MaxDist = temp.HitDist
            }
        } else {
            // Fork node
            var tMin, tMax Float
            if node.bbox.Intersect(r, &tMin, &tMax) {
                if node.left != nil {
                    pos += 1
                    stack[pos] = node.left
                }
                if node.right != nil {
                    pos += 1
                    stack[pos] = node.right
                }
            }
        }
    }
    return ret
}
