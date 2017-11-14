package gopt

import (
	"math/rand"
	"sort"

	. "github.com/tatsy/gopt/src/core"
)

const (
	BVH_SPLIT_METHOD_EQUAL_COUNT = iota
	BVH_SPLIT_METHOD_SAH
)

// -----------------------------------------------------------------------------
// BVH node
// -----------------------------------------------------------------------------
type BvhNode struct {
	left, right *BvhNode
	primId      int
	bbox        *Bounds3d
}

func NewLeafNode(primId int, b *Bounds3d) *BvhNode {
	node := &BvhNode{}
	node.left = nil
	node.right = nil
	node.primId = primId
	node.bbox = b
	return node
}

func NewForkNode(left *BvhNode, right *BvhNode, b *Bounds3d) *BvhNode {
	node := &BvhNode{}
	node.left = left
	node.right = right
	node.primId = -1
	node.bbox = b
	return node
}

func (node *BvhNode) IsLeaf() bool {
	return node.left == nil && node.right == nil
}

// -----------------------------------------------------------------------------
// BVH primitive info
// -----------------------------------------------------------------------------
type BvhPrimitiveInfo struct {
	primitiveId int
	centroid    *Vector3d
	bounds      *Bounds3d
}

func NewBvhPrimitiveInfo(primitiveId int, b *Bounds3d) *BvhPrimitiveInfo {
	info := &BvhPrimitiveInfo{}
	info.primitiveId = primitiveId
	info.bounds = b
	info.centroid = b.Centroid()
	return info
}

type Partitionable interface {
	sort.Interface
	Check(i int) bool
}

type BvhPrimitiveSortable struct {
	items   []*BvhPrimitiveInfo
	axis    int
	checker func(*BvhPrimitiveInfo) bool
}

func NewBvhPrimitiveSortable(items []*BvhPrimitiveInfo, axis int, checker func(*BvhPrimitiveInfo) bool) *BvhPrimitiveSortable {
	s := &BvhPrimitiveSortable{}
	s.items = items
	s.axis = axis
	s.checker = checker
	return s
}

func (s *BvhPrimitiveSortable) Len() int {
	return len(s.items)
}

func (s *BvhPrimitiveSortable) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s *BvhPrimitiveSortable) Less(i, j int) bool {
	return s.items[i].centroid.NthElement(s.axis) < s.items[j].centroid.NthElement(s.axis)
}

func (s *BvhPrimitiveSortable) Check(i int) bool {
	return s.checker(s.items[i])
}

// -----------------------------------------------------------------------------
// Bucket info
// -----------------------------------------------------------------------------
type BucketInfo struct {
	count  int
	bounds *Bounds3d
}

func NewBucketInfo(count int, bounds *Bounds3d) *BucketInfo {
	b := &BucketInfo{}
	b.count = count
	b.bounds = bounds
	return b
}

// -----------------------------------------------------------------------------
// Utility methods
// -----------------------------------------------------------------------------
func SlicePartition(items Partitionable) int {
	k := 0
	for i := 0; i < items.Len(); i++ {
		if items.Check(i) {
			if i != k {
				items.Swap(i, k)
			}
			k += 1
		}
	}
	return k
}

func SliceNthElement(items sort.Interface, start, end, N int) {
	if end-start <= 1 || end-start <= N || N == 0 {
		return
	} else if end-start == 2 {
		if !items.Less(start, start+1) {
			items.Swap(start, start+1)
		}
		return
	}

	pivot := start + rand.Intn(end-start)
	if pivot != end-1 {
		items.Swap(pivot, end-1)
		pivot = end - 1
	}

	k := start
	for i := start; i < end-1; i++ {
		if items.Less(i, pivot) {
			if i != k {
				items.Swap(i, k)
			}
			k += 1
		}
	}
	if k != pivot {
		items.Swap(k, pivot)
	}

	if (k - start) < N {
		SliceNthElement(items, k+1, end, N-(k-start+1))
	} else {
		SliceNthElement(items, start, k, N)
	}
}
