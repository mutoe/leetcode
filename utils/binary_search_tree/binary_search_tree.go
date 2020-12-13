package binary_search_tree

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

type BST struct {
	root *BinaryTree
	size int
}

func (this *BST) Add(num int) {
	this.size++
	this.root = add(this.root, num)
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
