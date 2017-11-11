package accelerator

import (
    "math"
    . "core"
)

type Bvh struct {
    primitives []*Primitive
    nodes []*BvhNode
    root *BvhNode
}

func NewBvh(primitives []*Primitive) *Bvh {
    bvh := &Bvh{}
    bvh.primitives = primitives

    buildData := make([]*BvhPrimitiveInfo, len(primitives))
    for i, p := range primitives {
        buildData[i] = NewBvhPrimitiveInfo(i, p.Bounds())
    }

    bvh.root = NewBvhSub(bvh, buildData)
    return bvh
}

func NewBvhSub(bvh *Bvh, buildData []*BvhPrimitiveInfo) *BvhNode {
    numData := len(buildData)
    if numData == 1 {
        node := NewLeafNode(buildData[0].primitiveId, buildData[0].bounds)
        bvh.nodes = append(bvh.nodes, node)
        return node
    }

    bounds := NewBounds3d()
    for i := 0; i < numData; i++ {
        bounds = bounds.Merge(buildData[i].bounds)
    }
    splitAxis := bounds.MaxExtent()

    splitMethod := BVH_SPLIT_METHOD_SAH
    splitPos := numData / 2

    switch splitMethod {
    case BVH_SPLIT_METHOD_EQUAL_COUNT:
        sortable := NewBvhPrimitiveSortable(
            buildData, splitAxis, nil,
        )
        SliceNthElement(sortable, 0, numData, splitPos)
    case BVH_SPLIT_METHOD_SAH:
        if numData <= 4 {
            sortable := NewBvhPrimitiveSortable(
                buildData, splitAxis, nil,
            )
            SliceNthElement(sortable, 0, numData, splitPos)
        } else {
            numBuckets := 12
            centroidBounds := NewBounds3d()
            for i := 0; i < numData; i++ {
                centroidBounds = centroidBounds.MergePoint(buildData[i].centroid)
            }

            buckets := make([]*BucketInfo, numBuckets)
            for i := 0; i < numBuckets; i++ {
                buckets[i] = NewBucketInfo(0, NewBounds3d())
            }

            cMin := centroidBounds.MinPos.NthElement(splitAxis)
            cMax := centroidBounds.MaxPos.NthElement(splitAxis)
            invDenom := 1.0 / (math.Abs(cMax - cMin) + Eps)
            for i := 0; i < numData; i++ {
                c0 := buildData[i].centroid.NthElement(splitAxis)
                c1 := centroidBounds.MinPos.NthElement(splitAxis)
                numer := c0 - c1
                b := int(Float(numBuckets) * math.Abs(numer) * invDenom)
                if b == numBuckets {
                    b = numBuckets - 1
                }

                buckets[b].count += 1
                buckets[b].bounds = buckets[b].bounds.Merge(buildData[i].bounds)
            }

            bucketCost := make([]Float, numBuckets - 1)
            for i := 0; i < numBuckets - 1; i++ {
                b0 := NewBounds3d()
                b1 := NewBounds3d()
                cnt0, cnt1 := 0, 0
                for j := 0; j <= i; j++ {
                    b0 = b0.Merge(buckets[j].bounds)
                    cnt0 += buckets[j].count
                }
                for j := i + 1; j < numBuckets; j++ {
                    b1 = b1.Merge(buckets[j].bounds)
                    cnt1 += buckets[j].count
                }
                bucketCost[i] += 0.125 + (Float(cnt0) * b0.Area() + Float(cnt1) * b1.Area()) / bounds.Area()
            }

            minCost := bucketCost[0]
            minCostSplit := 0
            for i := 1; i < numBuckets - 1; i++ {
                if minCost > bucketCost[i] {
                    minCost = bucketCost[i]
                    minCostSplit = i
                }
            }

            if minCost < Float(numData) {
                comparator := func(info *BvhPrimitiveInfo) bool {
                    cMin := centroidBounds.MinPos.NthElement(splitAxis)
                    cMax := centroidBounds.MaxPos.NthElement(splitAxis)
                    inv := 1.0 / (math.Abs(cMax - cMin) + Eps)
                    diff := math.Abs(info.centroid.NthElement(splitAxis) - cMin)
                    b := int(Float(numBuckets) * diff * inv)
                    if b >= numBuckets {
                        b = numBuckets - 1
                    }
                    return b <= minCostSplit
                }
                splitPos = SlicePartition(NewBvhPrimitiveSortable(
                    buildData, splitAxis, comparator,
                ))
            }
        }
    }

    leftNode := NewBvhSub(bvh, buildData[:splitPos])
    rightNode := NewBvhSub(bvh, buildData[splitPos:])
    node := NewForkNode(leftNode, rightNode, bounds)
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
            p := bvh.primitives[node.primId]
            if p.Intersect(r, &temp) {
                ret = true
                *isect = temp
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
