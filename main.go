package main

import "fmt"

type segmentTreeData struct {
	prefix int
	suffix int
	best   int
	total  int
}

func main() {
	s := createTree(1)
	fmt.Printf("%v", s)
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
