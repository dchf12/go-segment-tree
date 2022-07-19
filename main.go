package main

import (
	"fmt"
)

type segmentTreeData struct {
	prefix int
	suffix int
	best   int
	total  int
}
type segmentTree struct {
	n    int
	data []segmentTreeData
}

func main() {
	s := createData(1)
	fmt.Printf("%v", s)
	fmt.Printf("%v", merge(s, s))

}
func createData(v int) segmentTreeData {
	return segmentTreeData{
		prefix: v,
		suffix: v,
		best:   v,
		total:  v,
	}
}

func merge(left, right segmentTreeData) segmentTreeData {
	total := left.total + right.total
	prefix := max(left.prefix, left.total+right.prefix)
	suffix := max(right.suffix, right.total+left.suffix)
	best := max(
		left.best,
		right.best,
		prefix,
		suffix,
		left.suffix+right.prefix,
	)
	return segmentTreeData{
		prefix,
		suffix,
		best,
		total,
	}
}

func max(n ...int) int {
	var m int
	for _, v := range n {
		if v > m {
			m = v
		}
	}
	return m
}

func Build(arr []int) *segmentTree {
	n := len(arr)
	length := n * 4
	data := make([]segmentTreeData, length)
	tree := &segmentTree{
		n,
		data,
	}
	tree.build(arr, 1, 1, n)
	return tree
}

func (tree *segmentTree) build(arr []int, index, left, right int) {
	if left > right {
		return
	}
	if left == right {
		tree.data[index] = createData(arr[left-1])
	}
	if left < right {
		middle := (left + right) / 2
		tree.build(arr, index*2, left, middle)
		tree.build(arr, index*2+1, middle+1, right)
		tree.data[index] = merge(tree.data[index*2], tree.data[index*2+1])
	}
}
