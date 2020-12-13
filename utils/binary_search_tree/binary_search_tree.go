package binary_search_tree

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

type BST struct {
	root *BinaryTree
	size int
}

func (b *BST) Add(num int) {
	b.size++
	b.root = add(b.root, num)
}
func add(node *BinaryTree, num int) *BinaryTree {
	if node == nil {
		return &BinaryTree{Val: num}
	}
	if node.Val <= num {
		node.Right = add(node.Right, num)
	} else {
		node.Left = add(node.Left, num)
	}
	return node
}

func (b *BST) PrevOrder() {
	preOrder(b.root)
}
func preOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	preOrder(tree.Left)
	preOrder(tree.Right)
}

func (b *BST) MidOrder() {
	midOrder(b.root)
}
func midOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	midOrder(tree.Left)
	fmt.Println(tree.Val)
	midOrder(tree.Right)
}

func (b *BST) NextOrder() {
	nextOrder(b.root)
}
func nextOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	nextOrder(tree.Left)
	nextOrder(tree.Right)
	fmt.Println(tree.Val)
}

func (b *BST) LevelOrder() {
	queue := make([]*BinaryTree, 0)
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
