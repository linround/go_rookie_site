package main

import "fmt"

func main() {
	v := []int{1, 8, 3, 4, 5, 6, 7, 2, 9}
	sort(v)
	fmt.Println(v)
}

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func sort(values []int) {
	var root *tree
	// 构建一颗二叉树
	for _, v := range values {
		root = add(root, v)
	}
	// 前序遍历二叉树
	// 左中右
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
