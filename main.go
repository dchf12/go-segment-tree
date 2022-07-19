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
	s := createTree(1)
	fmt.Printf("%v", s)
	fmt.Printf("%v", merge(s, s))

}
func createTree(v int) segmentTreeData {
	return segmentTreeData{
		prefix: v,
		suffix: v,
		best:   v,
		total:  v,
	}
}

func merge(left, right segmentTreeData) segmentTreeData {
	total := left.total + right.total
	prefix := Max(left.prefix, left.total+right.prefix)
	suffix := Max(right.suffix, right.total+left.suffix)
	best := Max(
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

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
func Max(n ...int) int {
	var m int
	for _, v := range n {
		m = max(m, v)
	}
	return m
}
