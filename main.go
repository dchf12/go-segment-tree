package main

import "fmt"

type segmentTree struct {
	prefix int
	suffix int
	best   int
	total  int
}

func main() {
	s := createTree(1)
	fmt.Printf("%v", s)
}
func createTree(v int) segmentTree {
	return segmentTree{
		prefix: v,
		suffix: v,
		best:   v,
		total:  v,
	}
}
