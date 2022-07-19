package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func max(n ...int) int {
	var m int
	for _, v := range n {
		if v > m {
			m = v
		}
	}
	return m
}
func min(n ...int) int {
	m := 9999999999
	for _, v := range n {
		if v < m {
			m = v
		}
	}
	return m
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

func (tree *segmentTree) Update(x, y int) {
	tree.update(1, 1, tree.n, x, y)
}

func (tree *segmentTree) update(index, left, right, updateIndex, updateValue int) {
	if left > right || left > updateIndex || right < updateValue {
		return
	}
	if left == right {
		tree.data[index] = createData(updateValue)
	}
	if left < right {
		middle := (left + right) / 2

		tree.update(index*2, left, middle, updateIndex, updateValue)
		tree.update(index*2+1, middle+1, right, updateIndex, updateValue)
		tree.data[index] = merge(tree.data[index*2], tree.data[index*2+1])
	}
}

func (tree *segmentTree) Find(x, y int) int {
	return tree.find(1, 1, tree.n, x, y).best
}

func (tree *segmentTree) find(index int, left int, right int, findLeft int, findRight int) segmentTreeData {
	if left == findLeft && right == findRight {
		return tree.data[index]
	} else {
		middle := (left + right) / 2

		if findRight <= middle {
			return tree.find(index*2, left, middle, findLeft, findRight)
		} else if findLeft > middle {
			return tree.find(index*2+1, middle+1, right, findLeft, findRight)
		} else {
			leftResult := tree.find(index*2, left, middle, findLeft, min(middle, findRight))
			rightResult := tree.find(index*2+1, middle+1, right, max(findLeft, middle+1), findRight)
			mergedResult := merge(leftResult, rightResult)
			return mergedResult
		}
	}
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func readInt() int {
	var value int
	fmt.Fscanf(reader, "%d\n", &value)

	return value
}

func writeInt(value int) {
	fmt.Fprintln(writer, value)
}

func readArray(n int) []int {
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	stringArray := strings.Split(strings.TrimSpace(line), " ")
	if len(stringArray) != n {
		panic(fmt.Errorf("Expected input array to be of size %d, but was %d", n, len(stringArray)))
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		value, err := strconv.Atoi(stringArray[i])
		if err != nil {
			panic(err)
		}

		arr[i] = value
	}

	return arr
}

func main() {
	defer writer.Flush()

	n := readInt()
	arr := readArray(n)
	tree := Build(arr)

	m := readInt()
	for i := 0; i < m; i++ {
		query := readArray(3)
		t := query[0]
		x := query[1]
		y := query[2]

		if t == 0 {
			tree.Update(x, y)
		} else {
			value := tree.Find(x, y)
			writeInt(value)
		}
	}
}
