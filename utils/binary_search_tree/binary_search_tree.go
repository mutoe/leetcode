package binary_search_tree

import (
	"fmt"

	. "../tree_node"
)

type BinarySearchTree struct {
	root *TreeNode
	size int
}

func (b *BinarySearchTree) Add(num int) {
	b.size++
	b.root = add(b.root, num)
}
func add(node *TreeNode, num int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: num}
	}
	if node.Val <= num {
		node.Right = add(node.Right, num)
	} else {
		node.Left = add(node.Left, num)
	}
	return node
}

func (b *BinarySearchTree) PrevOrder() {
	preOrder(b.root)
}
func preOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	preOrder(tree.Left)
	preOrder(tree.Right)
}

func (b *BinarySearchTree) MidOrder() {
	midOrder(b.root)
}
func midOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	midOrder(tree.Left)
	fmt.Println(tree.Val)
	midOrder(tree.Right)
}

func (b *BinarySearchTree) NextOrder() {
	nextOrder(b.root)
}
func nextOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	nextOrder(tree.Left)
	nextOrder(tree.Right)
	fmt.Println(tree.Val)
}

func (b *BinarySearchTree) LevelOrder() {
	queue := make([]*TreeNode, 0)
	queue = append(queue, b.root)
	for len(queue) > 0 {
		cur := queue[0]
		fmt.Println(cur.Val)
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}
