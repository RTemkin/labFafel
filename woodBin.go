package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	len  int
}

func (b *bst) insert(value int) {
	newNode := &node{value: value}

	if b.root == nil {
		b.root = newNode
	} else {
		inertNode(b.root, newNode)
	}
}

func inertNode(current, newNode *node) {
	if newNode.value < current.value {
		if current.left == nil {
			current.left = newNode
		} else {
			inertNode(current.left, newNode)
		}
	} else {
		if current.right == nil {
			current.right = newNode
		} else {
			inertNode(current.right, newNode)
		}
	}
}

func (b *bst) InOrderTreversal(n *node) {
	if n != nil {
		b.InOrderTreversal(n.left)
		fmt.Printf("%d", n.value)
		b.InOrderTreversal(n.right)
	}
}

func main() {

	tree := bst{}

	values := []int{5, 2, 8, 1, 3, 7, 9}

	for _, value := range values {
		tree.insert(value)
	}

	tree.InOrderTreversal(tree.root)
}
